package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 假设你有一个结构体来存储请求数据
type VoiceParams struct {
	VoiceType   string `json:"voice_type"`
	Emotion     string `json:"emotion"`
	Lang        string `json:"lang"`
	AudioFormat string `json:"audio_format"`
	Rate        string `json:"rate"`
	Pitch       string `json:"pitch"`
	Volume      string `json:"volume"`
}

// 假设你有一个结构体来存储响应数据
type ResponseData struct {
	Data string `json:"data"`
}

// 假设你有一个函数来从文本获取音频数据
//func fetchAudioFromText(apiURL, text, traceID string, voiceParams *VoiceParams) (data []byte, timeLog string, err error) {
//	// 构建请求数据
//	requestData := map[string]interface{}{
//		"text": text,
//		"voice_params": map[string]interface{}{
//			"voice_type":   voiceParams.VoiceType,
//			"emotion":      voiceParams.Emotion,
//			"lang":         voiceParams.Lang,
//			"audio_format": voiceParams.AudioFormat,
//			"rate":         voiceParams.Rate,
//			"pitch":        voiceParams.Pitch,
//			"volume":       voiceParams.Volume,
//		},
//		"trace_id": traceID,
//	}
//
//	// 发送 POST 请求到服务器
//	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer([]byte(json.Marshal(requestData))))
//	if err != nil {
//		return nil, "", err
//	}
//	defer response.Body.Close()
//
//	if response.StatusCode != 200 {
//		return nil, "", fmt.Errorf("Failed to fetch audio data: %d", response.StatusCode)
//	}
//
//	// 解析响应数据
//	responseData := ResponseData{}
//	err = json.NewDecoder(response.Body).Decode(&responseData)
//	if err != nil {
//		return nil, "", err
//	}
//
//	data, err = base64.StdEncoding.DecodeString(responseData.Data)
//	if err != nil {
//		return nil, "", err
//	}
//	timeLog = responseData.TimeLog
//
//	return data, timeLog, nil
//}

// 假设你有一个函数来从SSML获取音频数据
func fetchAudioFromSSML(apiURL, ssmlFile, traceID string) (data []byte, timeLog string, err error) {
	// 发送 POST 请求到服务器
	//ssmlFile = "<speak voice_type=\"xiaosi\" rate=\"1\" pitch=\"1\" volume=\"1\" emotion='confused'>\n    <break time=\"3000ms\"/>\n    屋漏偏逢连夜雨<break time=\"3000ms\"/>\n    船迟又遇打头风<break time=\"3000ms\"/>\n</speak>"
	param := map[string]string{
		"text": "<speak voice_type=\"xiaosi\" rate=\"1\" pitch=\"1\" volume=\"1\" emotion='confused'>\n    屋漏偏逢连夜雨\n</speak>",
	}
	marshal, err := json.Marshal(param)
	if err != nil {
		return
	}
	response, err := http.Post(apiURL, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, "", fmt.Errorf("Failed to fetch audio data: %d", response.StatusCode)
	}
	// 解析响应数据
	responseData := ResponseData{}
	//json.Unmarshal(marshal, &responseData)
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		return nil, "", err
	}

	// 解析响应数据
	//base64转码
	fmt.Println(responseData.Data)
	decodeString, err := base64.StdEncoding.DecodeString(responseData.Data)
	if err != nil {
		return nil, "", err
	}
	err = ioutil.WriteFile(fmt.Sprintf("output_111.mp3"), decodeString, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func main5() {
	fetchAudioFromSSML("http://hmi.chengjiukehu.com/tal-vits-pipeline", "", "")

}

func main() {

	url := "http://hmi-in.chengjiukehu.com/tal-vits-pipeline"
	method := "POST"
	//temp := VoiceParams{
	//	VoiceType:   "xiaosi",
	//	Emotion:     "happy",
	//	Lang:        "cn",
	//	AudioFormat: "mp3",
	//	Rate:        "1",
	//	Pitch:       "1",
	//	Volume:      "1",
	//}
	//ma, err := json.Marshal(temp)
	param := map[string]string{
		"text": "<speak voice_type=\"xiaosi\" rate=\"1\" pitch=\"1\" volume=\"100\" emotion=\"happy\"> <phoneme lang=\"cn\" ph=\"yi4/huir2/jia1/ban3\">一会儿加班。</phoneme></speak>",
		//"text":         "你好",
		//"voice_params": string(ma),
	}
	//a := "<speak voice_type=\"xiaosi\" emotion=\"happy\" volume=\"0.9\" rate=\"1.1\" pitch=\"1.0\" lang=\"cn\">\n<break time=\\\"100ms\\\"/>\n你好啊,今天天气\n</speak>\n"

	marshal, err := json.Marshal(param)
	if err != nil {
		return
	}
	payload := strings.NewReader(string(marshal))
	//	payload = strings.NewReader(`{
	//    "text": "<speak voice_type=\"xiaosi\" rate=\"1\" pitch=\"1\" volume=\"1\" emotion='confused'>\n    <break time=\"3000ms\"/>\n    屋漏偏逢连夜雨<break time=\"3000ms\"/>\n    船迟又遇打头风<break time=\"3000ms\"/>\n</speak>",
	//    "trace_id":"1"
	//}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var a ResponseData
	err = json.Unmarshal(body, &a)
	if err != nil {
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(a.Data)
	err = ioutil.WriteFile(fmt.Sprintf("output_123.mp3"), decodeString, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(body))
}
