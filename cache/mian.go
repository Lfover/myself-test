package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"math"
	"sync"
	"time"
)

// TTSRedisCache Redis缓存管理器
type TTSRedisCache struct {
	client     *redis.Client
	maxItems   int
	updateLock sync.Mutex
}

// NewTTSRedisCache 创建缓存管理器
func NewTTSRedisCache(redisAddr string, maxItems int) *TTSRedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // 设置密码
		DB:       0,  // 使用默认DB
	})

	cache := &TTSRedisCache{
		client:   client,
		maxItems: maxItems,
	}

	// 启动维护协程
	go cache.startMaintenanceTask()

	return cache
}

// 记录访问并更新热度分数
func recordAccess(hash string, redisClient *redis.Client) {
	now := time.Now()
	pipe := redisClient.Pipeline()

	// 1. 递增访问计数
	pipe.Incr(fmt.Sprintf("tts:access:%s", hash))

	// 2. 更新最后访问时间
	metadata, _ := pipe.HGet("tts:cache:metadata", hash).Result()
	if metadata != "" {
		metaObj := parseMetadata(metadata)
		metaObj.LastAccess = now
		metaObj.AccessCount++

		// 更新5分钟、1小时、24小时时间窗口计数
		updateTimeWindows(metaObj, now)

		pipe.HSet("tts:cache:metadata", hash, serializeMetadata(metaObj))
	}

	// 3. 重新计算热度分数
	score := calculateHotScore(hash, redisClient)
	pipe.ZAdd("tts:cache:hotrank", &redis.Z{
		Score:  score,
		Member: hash,
	})

	pipe.Exec()
}

// 更新时间窗口计数
func updateTimeWindows(meta *Metadata, now time.Time) {
	// 5分钟时间窗口计数
	if now.Sub(meta.LastWindowReset5m) > 5*time.Minute {
		meta.CountWindow5m = 1
		meta.LastWindowReset5m = now
	} else {
		meta.CountWindow5m++
	}

	// 1小时时间窗口
	if now.Sub(meta.LastWindowReset1h) > time.Hour {
		meta.CountWindow1h = 1
		meta.LastWindowReset1h = now
	} else {
		meta.CountWindow1h++
	}

	// 24小时时间窗口
	if now.Sub(meta.LastWindowReset24h) > 24*time.Hour {
		meta.CountWindow24h = 1
		meta.LastWindowReset24h = now
	} else {
		meta.CountWindow24h++
	}
}

// 判断是否应该放入缓存
func shouldCacheInRedis(hash string, redisClient *redis.Client) bool {
	// 检查缓存是否已达上限
	cacheSize, _ := redisClient.ZCard("tts:cache:hotrank").Result()

	// 如果缓存未满，直接放入
	if cacheSize < 1000 {
		return true
	}

	// 缓存已满，只有比现有最冷的缓存项更热才放入
	lowestScore, _ := getLowestHotScore(redisClient)
	currentScore := calculateHotScore(hash, redisClient)

	// 设置一个动态准入阈值(比最低分数高出20%)
	admissionThreshold := lowestScore * 1.2

	return currentScore > admissionThreshold
}

// 获取当前缓存中最低热度分数
func getLowestHotScore(redisClient *redis.Client) (float64, error) {
	result, err := redisClient.ZRange("tts:cache:hotrank", 0, 0).Result()
	if err != nil || len(result) == 0 {
		return 0, err
	}

	score, _ := redisClient.ZScore("tts:cache:hotrank", result[0]).Result()
	return score, nil
}

// 定期更新热度分数(每小时执行)
func periodicHotnessUpdate(redisClient *redis.Client) {
	// 获取所有缓存项
	allItems, _ := redisClient.ZRange("tts:cache:hotrank", 0, -1).Result()

	pipe := redisClient.Pipeline()
	for _, hash := range allItems {
		// 重新计算热度分数
		score := calculateHotScore(hash, redisClient)

		// 更新Sorted Set
		pipe.ZAdd("tts:cache:hotrank", &redis.Z{
			Score:  score,
			Member: hash,
		})
	}
	pipe.Exec()

	// 清理低于阈值的项
	purgeColditems(redisClient)
}

// 清理冷数据
func purgeColditems(redisClient *redis.Client) {
	// 获取当前缓存大小
	cacheSize, _ := redisClient.ZCard("tts:cache:hotrank").Result()

	// 如果缓存未达到淘汰阈值，不执行淘汰
	if cacheSize <= 900 { // 预留100个缓冲区
		return
	}

	// 计算需要淘汰的数量(保留900项)
	toPurge := int(cacheSize - 900)

	// 获取热度最低的N个项
	coldItems, _ := redisClient.ZRange("tts:cache:hotrank", 0, int64(toPurge-1)).Result()

	pipe := redisClient.Pipeline()
	for _, hash := range coldItems {
		// 从热度排行中移除
		pipe.ZRem("tts:cache:hotrank", hash)

		// 保留映射和元数据(可选，取决于内存压力)
		// 仅当内存紧张时才彻底删除
		if shouldRemoveCompletely() {
			pipe.HDel("tts:cache:mapping", hash)
			pipe.HDel("tts:cache:metadata", hash)
			pipe.Del(fmt.Sprintf("tts:access:%s", hash))
		}
	}
	pipe.Exec()
}

// 是否应该完全删除缓存项(可基于Redis内存使用率判断)
func shouldRemoveCompletely() bool {
	// 此处可判断Redis内存使用情况
	// 例如通过INFO MEMORY命令获取使用率
	return false // 默认保留映射和元数据
}

// 计算热度分数
func calculateHotScore(hash string, redisClient *redis.Client) float64 {
	// 基础计数 - 请求次数
	count, _ := redisClient.Get(fmt.Sprintf("tts:access:%s", hash)).Int64()

	// 获取元数据
	metadata, _ := redisClient.HGet("tts:cache:metadata", hash).Result()
	metaObj := parseMetadata(metadata)

	// 计算时间衰减因子
	timeDecay := calculateTimeDecay(metaObj.LastAccess)

	// 计算热度分数 (基于时间衰减的加权访问频率)
	baseScore := float64(count)

	// 应用时间衰减
	// 一天内保持100%，之后每天衰减30%
	score := baseScore * timeDecay

	// 额外因素调整
	if metaObj.TextLength < 10 {
		// 超短文本加权(通常是高频短句)
		score *= 1.2
	}

	return score
}

// 计算时间衰减因子
func calculateTimeDecay(lastAccess time.Time) float64 {
	hoursSince := time.Since(lastAccess).Hours()

	if hoursSince < 24 {
		return 1.0 // 24小时内不衰减
	} else {
		// 每24小时衰减30%
		daysSince := hoursSince / 24
		return math.Pow(0.7, daysSince)
	}
}

// 获取音频URL
func (c *TTSRedisCache) GetAudioURL(text string, params TTSParams) (string, bool) {
	hash := generateCacheKey(text, params)

	// 从Redis获取URL
	url, err := c.client.HGet("tts:cache:mapping", hash).Result()
	if err != nil || url == "" {
		return "", false
	}

	// 异步更新访问统计
	go c.recordAccess(hash)

	return url, true
}

// 缓存音频URL
func (c *TTSRedisCache) CacheAudioURL(text string, params TTSParams, url string) {
	hash := generateCacheKey(text, params)

	// 判断是否值得缓存
	if !shouldCacheInRedis(hash, c.client) {
		return
	}

	pipe := c.client.Pipeline()

	// 存储URL映射
	pipe.HSet("tts:cache:mapping", hash, url)

	// 存储元数据
	metadata := Metadata{
		Text:               text,
		Params:             params,
		CreatedAt:          time.Now(),
		LastAccess:         time.Now(),
		AccessCount:        1,
		CountWindow5m:      1,
		CountWindow1h:      1,
		CountWindow24h:     1,
		LastWindowReset5m:  time.Now(),
		LastWindowReset1h:  time.Now(),
		LastWindowReset24h: time.Now(),
	}
	pipe.HSet("tts:cache:metadata", hash, serializeMetadata(metadata))

	// 初始化访问计数
	pipe.Set(fmt.Sprintf("tts:access:%s", hash), 1, 0)

	// 计算初始热度分数
	initialScore := 1.0
	pipe.ZAdd("tts:cache:hotrank", &redis.Z{
		Score:  initialScore,
		Member: hash,
	})

	pipe.Exec()
}

// 启动维护任务
func (c *TTSRedisCache) startMaintenanceTask() {
	hourlyTicker := time.NewTicker(1 * time.Hour)
	dailyTicker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-hourlyTicker.C:
			// 执行热度分数更新
			c.updateLock.Lock()
			periodicHotnessUpdate(c.client)
			c.updateLock.Unlock()

		case <-dailyTicker.C:
			// 执行深度清理
			c.updateLock.Lock()
			c.deepCleanup()
			c.updateLock.Unlock()
		}
	}
}

// TTSService 服务层
type TTSService struct {
	redisCache  *TTSRedisCache
	mysqlRepo   *TTSMySQLRepository
	ttsProvider *TTSProvider
}

// GetAudio 获取音频
func (s *TTSService) GetAudio(text string, params TTSParams) (string, error) {
	// 1. 尝试从Redis缓存获取
	url, found := s.redisCache.GetAudioURL(text, params)
	if found {
		return url, nil
	}

	// 2. Redis未命中，尝试从MySQL获取
	hash := generateCacheKey(text, params)
	url, found = s.mysqlRepo.GetURLByHash(hash)

	if found {
		// 找到了URL，异步添加到Redis缓存
		go s.redisCache.CacheAudioURL(text, params, url)
		return url, nil
	}

	// 3. MySQL也未命中，调用TTS服务生成
	url, err := s.ttsProvider.SynthesizeSpeech(text, params)
	if err != nil {
		return "", err
	}

	// 4. 同时保存到MySQL和Redis
	go func() {
		s.mysqlRepo.SaveURL(hash, text, params, url)
		s.redisCache.CacheAudioURL(text, params, url)
	}()

	return url, nil
}

// TTSService 服务层
type TTSService struct {
	redisCache  *TTSRedisCache
	mysqlRepo   *TTSMySQLRepository
	ttsProvider *TTSProvider
}

// GetAudio 获取音频
func (s *TTSService) GetAudio(text string, params TTSParams) (string, error) {
	// 1. 尝试从Redis缓存获取
	url, found := s.redisCache.GetAudioURL(text, params)
	if found {
		return url, nil
	}

	// 2. Redis未命中，尝试从MySQL获取
	hash := generateCacheKey(text, params)
	url, found = s.mysqlRepo.GetURLByHash(hash)

	if found {
		// 找到了URL，异步添加到Redis缓存
		go s.redisCache.CacheAudioURL(text, params, url)
		return url, nil
	}

	// 3. MySQL也未命中，调用TTS服务生成
	url, err := s.ttsProvider.SynthesizeSpeech(text, params)
	if err != nil {
		return "", err
	}

	// 4. 同时保存到MySQL和Redis
	go func() {
		s.mysqlRepo.SaveURL(hash, text, params, url)
		s.redisCache.CacheAudioURL(text, params, url)
	}()

	return url, nil
}

// 系统启动时预热缓存
func (c *TTSRedisCache) preheatCache() {
	// 从MySQL加载最近一段时间内最常访问的1000条记录
	topItems, _ := mysqlRepo.GetTopAccessedItems(1000)

	pipe := c.client.Pipeline()
	for _, item := range topItems {
		// 添加到Redis缓存
		pipe.HSet("tts:cache:mapping", item.Hash, item.URL)

		// 设置初始热度分数(基于历史访问频率)
		initialScore := float64(item.AccessCount)
		pipe.ZAdd("tts:cache:hotrank", &redis.Z{
			Score:  initialScore,
			Member: item.Hash,
		})

		// 初始化访问计数和元数据
		pipe.Set(fmt.Sprintf("tts:access:%s", item.Hash), item.AccessCount, 0)

		metadata := Metadata{
			Text:      item.Text,
			Params:    deserializeParams(item.Params),
			CreatedAt: item.CreatedAt,
			// 其他字段初始化...
		}
		pipe.HSet("tts:cache:metadata", item.Hash, serializeMetadata(metadata))
	}
	pipe.Exec()
}

// 导出Redis缓存性能指标
func (c *TTSRedisCache) ExportMetrics() map[string]interface{} {
	// 获取缓存大小
	cacheSize, _ := c.client.ZCard("tts:cache:hotrank").Result()

	// 获取热点分布情况
	scores := []float64{}
	topItems, _ := c.client.ZRevRangeWithScores("tts:cache:hotrank", 0, 9).Result()
	for _, z := range topItems {
		scores = append(scores, z.Score)
	}

	// 返回监控指标
	return map[string]interface{}{
		"cache_size":       cacheSize,
		"hit_rate":         c.hitRate,
		"top10_scores":     scores,
		"memory_usage_kb":  getRedisMemoryUsage(c.client),
		"cache_operations": c.operationCounter,
	}
}
