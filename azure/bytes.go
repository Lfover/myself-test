package main

import (
	"fmt"
)

// 定义字边界事件结构体
type WordBoundary struct {
	BoundaryType string // Word 或 Punctuation
	AudioOffset  int64  // 音频偏移量(ms)
	Duration     int64  // 持续时间(ns)
	Text         string // 文本内容
	TextOffset   int    // 文本偏移量
	WordLength   int    // 文本长度
}

// 用于追踪音频流进度的结构体
type AudioProgress struct {
	totalBytes int64          // 音频总字节数
	boundaries []WordBoundary // 所有字边界事件
	bytesPerMs float64        // 每毫秒对应的字节数
}

// 创建新的 AudioProgress
func NewAudioProgress(totalBytes int64, boundaries []WordBoundary) *AudioProgress {
	// 获取最后一个边界的时间作为总时长
	lastBoundary := boundaries[len(boundaries)-1]
	totalMs := float64(lastBoundary.AudioOffset + (lastBoundary.Duration / 1000000)) // 转换为毫秒
	bytesPerMs := float64(totalBytes) / totalMs

	return &AudioProgress{
		totalBytes: totalBytes,
		boundaries: boundaries,
		bytesPerMs: bytesPerMs,
	}
}

// 根据当前字节数获取对应的单词信息
func (ap *AudioProgress) GetCurrentWord(currentBytes int64) *WordBoundary {
	// 计算当前大约处于哪个时间点(ms)
	currentMs := float64(currentBytes) / ap.bytesPerMs

	// 查找对应的边界事件
	var currentBoundary *WordBoundary
	for i, boundary := range ap.boundaries {
		// 计算该边界的结束时间
		boundaryEndMs := float64(boundary.AudioOffset + (boundary.Duration / 1000000))

		if float64(boundary.AudioOffset) <= currentMs && currentMs <= boundaryEndMs {
			return &ap.boundaries[i]
		}

		// 如果当前时间在两个边界之间，返回前一个边界
		if i < len(ap.boundaries)-1 {
			nextBoundaryStartMs := float64(ap.boundaries[i+1].AudioOffset)
			if currentMs > boundaryEndMs && currentMs < nextBoundaryStartMs {
				return &ap.boundaries[i]
			}
		}
	}

	return currentBoundary
}

// 使用示例
func main2() {
	// 示例数据
	boundaries := []WordBoundary{
		{
			BoundaryType: "Word",
			AudioOffset:  500000,
			Duration:     175000000,
			Text:         "The",
			TextOffset:   250,
			WordLength:   3,
		},
		{
			BoundaryType: "Word",
			AudioOffset:  2375000,
			Duration:     425000000,
			Text:         "rainbow",
			TextOffset:   254,
			WordLength:   7,
		},
		{
			BoundaryType: "Punctuation",
			AudioOffset:  79125000,
			Duration:     87500000,
			Text:         "..",
			TextOffset:   4294967295,
			WordLength:   2,
		},
	}

	// 假设总字节数为 10000
	totalBytes := int64(268846)
	progress := NewAudioProgress(totalBytes, boundaries)

	// 在流式处理中使用
	chunkSize := int64(12000)
	var processedBytes int64

	for processedBytes < totalBytes {
		// 模拟处理一个数据块
		processedBytes += chunkSize

		// 获取当前对应的单词
		if word := progress.GetCurrentWord(processedBytes); word != nil {
			fmt.Printf("当前处理至字节: %d, 对应文本: %s, 时间点: %dms\n",
				processedBytes, word.Text, word.AudioOffset)
		}
	}
}

func main() {
	boundaries := []WordBoundary{
		{
			BoundaryType: "Punctuation",
			AudioOffset:  625000,
			Duration:     531250000,
			Text:         "你好",
			TextOffset:   244,
			WordLength:   2,
		},
		{
			BoundaryType: "Punctuation",
			AudioOffset:  5937500,
			Duration:     390625000,
			Text:         "ya",
			TextOffset:   246,
			WordLength:   2,
		},
	}
	temp := GetAudioProgress(60750, boundaries)
	temp.totalBytes = 60750
	temp.boundaries = boundaries
	temp.bytesPerMs = 0.010230906938323209
	var i int
	for {
		if 60750-i >= 12000 {
			i += 12000
		} else {
			i += 60750 - i - 1
		}
		if i >= 60750 {
			break
		}
		te := GetCurrentWord(int64(i), temp)
		fmt.Println(te.Text)
		fmt.Println(temp)
		fmt.Println(te)
	}

}
func GetAudioProgress(totalBytes int, boundaries []WordBoundary) *AudioProgress {
	// 获取最后一个边界的时间作为总时长
	if len(boundaries) == 0 {
		fmt.Println("No boundaries found")
		return nil
	}
	lastBoundary := boundaries[len(boundaries)-1]
	totalMs := float64(lastBoundary.AudioOffset + (lastBoundary.Duration / 1000000)) // 转换为毫秒
	bytesPerMs := float64(totalBytes) / totalMs
	//fmt.Printf("bytesPerMs:%+v,lastBoundary:%+v,totalBytes:%+v", bytesPerMs, lastBoundary, totalBytes)
	return &AudioProgress{
		totalBytes: int64(totalBytes),
		boundaries: boundaries,
		bytesPerMs: bytesPerMs,
	}
}
func GetCurrentWord(currentBytes int64, ap *AudioProgress) *WordBoundary {
	// 计算当前大约处于哪个时间点(ms)
	currentMs := float64(currentBytes) / ap.bytesPerMs

	// 查找对应的边界事件
	var currentBoundary *WordBoundary
	for i, boundary := range ap.boundaries {
		// 计算该边界的结束时间
		boundaryEndMs := float64(boundary.AudioOffset + (boundary.Duration / 1000000))

		if float64(boundary.AudioOffset) <= currentMs && currentMs <= boundaryEndMs {
			return &ap.boundaries[i]
		}

		// 如果当前时间在两个边界之间，返回前一个边界
		if i < len(ap.boundaries)-1 {
			nextBoundaryStartMs := float64(ap.boundaries[i+1].AudioOffset)
			if currentMs > boundaryEndMs && currentMs < nextBoundaryStartMs {
				return &ap.boundaries[i]
			}
		}
	}

	return currentBoundary
}
