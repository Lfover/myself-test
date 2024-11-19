package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"

	"golang.org/x/sound"
)

// 假设你有一个Base64字符串

func main() {
	// 解码Base64字符串
	pcmData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal(err)
	}

	// 创建PCM数据切片
	pcm := sound.NewPCM(sound.FormatS16, 1, 44100)

	// 填充PCM数据
	err = pcm.Load(pcmData)
	if err != nil {
		log.Fatal(err)
	}

	// 创建MP3编码器
	mp3 := sound.NewMP3Encoder(44100)

	// 创建输出缓冲区
	output := &bytes.Buffer{}

	// 开始编码
	err = mp3.Encode(output, pcm)
	if err != nil {
		log.Fatal(err)
	}

	// 保存MP3文件
	err = ioutil.WriteFile("output1111.mp3", output.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PCM数据已成功转换为MP3文件")
}
