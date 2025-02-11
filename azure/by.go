package main

import "fmt"

type AudioTextTracker struct {
	boundaries    []WordBoundary
	sampleRate    int     // 采样率，例如 16000
	bytesPerMs    float64 // 每毫秒的字节数
	bytesPerChunk int     // 每个数据包的大小
}

func NewAudioTextTracker(boundaries []WordBoundary, sampleRate int, bytesPerChunk int) *AudioTextTracker {
	// 16位音频，单声道，则每毫秒的字节数 = (采样率 * 2) / 1000
	bytesPerMs := float64(sampleRate*2) / 1000.0

	return &AudioTextTracker{
		boundaries:    boundaries,
		sampleRate:    sampleRate,
		bytesPerMs:    bytesPerMs,
		bytesPerChunk: bytesPerChunk,
	}
}

// 获取指定数据包对应的文本信息
func (t *AudioTextTracker) GetTextForChunk(chunkIndex int) []WordBoundary {
	// 计算当前数据包对应的时间范围
	chunkStartBytes := int64(chunkIndex * t.bytesPerChunk)
	chunkEndBytes := int64((chunkIndex + 1) * t.bytesPerChunk)

	// 转换为毫秒时间
	chunkStartMs := float64(chunkStartBytes) / t.bytesPerMs
	chunkEndMs := float64(chunkEndBytes) / t.bytesPerMs

	// 查找这个时间范围内的所有文本
	var result []WordBoundary
	for _, boundary := range t.boundaries {
		// 计算每个边界的时间范围
		wordStartMs := float64(boundary.AudioOffset)
		wordEndMs := float64(boundary.AudioOffset) + float64(boundary.Duration)/1000000.0

		// 检查是否有重叠
		if (wordStartMs <= chunkEndMs) && (wordEndMs >= chunkStartMs) {
			result = append(result, boundary)
		}
	}

	return result
}

// 使用示例
func mai1n() {
	// 示例边界数据
	boundaries := []WordBoundary{
		{
			BoundaryType: "Word",
			AudioOffset:  3750,
			Duration:     475000000,
			Text:         "yellow",
			TextOffset:   329,
			WordLength:   6,
		},
		{
			BoundaryType: "Punctuation",
			AudioOffset:  4325,
			Duration:     162500000,
			Text:         ",",
			TextOffset:   335,
			WordLength:   1,
		},
		// ... 其他边界数据 ...
	}

	// 创建追踪器（假设采样率为16000，每个数据包1200字节）
	tracker := NewAudioTextTracker(boundaries, 16000, 1200)

	// 在流式处理中使用
	var chunkIndex int
	for {
		// 获取当前数据包对应的文本
		words := tracker.GetTextForChunk(chunkIndex)
		if len(words) > 0 {
			fmt.Printf("Chunk %d (%d-%d ms) 包含以下文本:\n",
				chunkIndex,
				int(float64(chunkIndex*1200)/tracker.bytesPerMs),
				int(float64((chunkIndex+1)*1200)/tracker.bytesPerMs))

			for _, word := range words {
				fmt.Printf("  - %s (offset: %dms, duration: %dms)\n",
					word.Text,
					word.AudioOffset,
					word.Duration/1000000)
			}
		}

		chunkIndex++
		// ... 处理音频数据 ...
	}
}
