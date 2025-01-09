package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 创建 Gin 实例
	r := gin.Default()

	// 配置 CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"POST", "GET", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	// 创建临时文件夹
	if err := os.MkdirAll("temp", os.ModePerm); err != nil {
		fmt.Printf("创建临时文件夹失败: %v\n", err)
		return
	}

	// 处理音频转换的接口
	r.POST("/convert", func(c *gin.Context) {
		file, err := c.FormFile("audio")
		if err != nil {
			c.JSON(400, gin.H{
				"error": fmt.Sprintf("文件上传失败: %v", err),
			})
			return
		}

		// 生成唯一文件名
		filename := uuid.New().String()
		pcmPath := filepath.Join("temp", filename+".pcm")
		mp3Path := filepath.Join("temp", filename+".mp3")

		// 保存上传的PCM文件
		if err := c.SaveUploadedFile(file, pcmPath); err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("文件保存失败: %v", err),
			})
			return
		}

		// 使用FFmpeg转换音频
		cmd := exec.Command("ffmpeg",
			"-f", "s16le", // PCM格式
			"-ar", "24000", // 采样率
			"-ac", "1", // 单声道
			"-i", pcmPath, // 输入文件
			"-acodec", "libmp3lame", // MP3编码器
			mp3Path) // 输出文件

		output, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("音频转换失败: %v, 输出: %s", err, string(output)),
			})
			return
		}

		// 返回转换后的文件
		c.File(mp3Path)

		// 清理临时文件
		defer func() {
			os.Remove(pcmPath)
			os.Remove(mp3Path)
		}()
	})

	// 添加健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	fmt.Println("服务器启动在 http://localhost:8009")
	if err := r.Run(":8009"); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
