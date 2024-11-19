package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type TtsBiz uint8
type TtsSource uint8

type Sentence struct {
	VoiceName     string
	ExpressStyle  string
	ProsodyVolume string
	ProsodyRate   string
	ProsodyPitch  string
	Text          string
}
type DataS struct {
	Sentences []Sentence
	Lang      string
	VoiceName string
}
type TtsRequest struct {
	Biz       TtsBiz     `json:"biz" form:"biz" binding:"required"`
	Source    TtsSource  `json:"source" form:"source" binding:"required"`
	Text      string     `json:"text" form:"text" binding:"required"`
	Phoneme   TtsPhoneme `json:"phoneme" form:"phoneme"`
	Lang      string     `json:"lang" form:"lang" binding:"required"` //默认语言,zh-CN, en-US
	Params    TtsParams  `json:"params" form:"params"`
	Ext       TtsExt     `json:"ext" form:"ext"`
	ChunkSize int        `json:"chunk_size" form:"chunk_size"`
	IsWav     uint8      `json:"is_wav" form:"is_wav"` //是否生成wav
	IsDb      uint8      `json:"is_db" form:"is_db"`   //是否存入数据库
}
type TtsPhoneme struct {
	Text string `json:"text"`
	Ph   string `json:"ph"`
}
type TtsParams struct {
	Voice      string `json:"voice"`       //语音名称
	Style      string `json:"style"`       //讲话风格
	Role       string `json:"role"`        //讲话角色扮演
	Pitch      string `json:"pitch"`       //音调
	Rate       string `json:"rate"`        //速率
	Volume     string `json:"volume"`      //音量
	SampleRate int    `json:"sample_rate"` //采样率
}
type TtsExt struct {
	SSML string `json:"ssml"`
}
type TtsResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	TraceId   string `json:"trace_id"`
	IsEnd     bool   `json:"is_end"`
	Result    []byte `json:"result"`
}

var (
	endpoint        = "https://southeastasia.voice.speech.microsoft.com/cognitiveservices/v1?deploymentId=7b7e1142-cf9a-4ee5-a0cd-ee14593886f9"
	subscriptionKey = "11af1a1a48f844838746d09dd27efeb0"
)

func main2() {
	req := TtsRequest{
		Biz:    2,
		Source: 1,
		Text:   "",
		Lang:   "zh-CN",
		Params: TtsParams{
			Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
			Pitch:  "+5Hz",
			Rate:   "1.1",
			Volume: "90",
		},
	}
	a := []map[string]string{
		{
			"你好我要说一件严肃的事情":  "serious",
			"你有什么技能?":       "cheerful",
			"他后来被亳州刺史闾丘晓所杀": "sad",
		},
		{
			"你好，我要说一件严肃的事a": "chat",
			"你有什么技能a?":      "cheerful",
		},
		{
			"王昌龄（698年－756年），字少伯，是唐代著名的边塞诗人。他在唐玄宗开元十五年（727年）考中进士，": "chat",
			"他又通过了博学宏词科的考试。": "cheerful",
		},
	}
	for i, text := range a {
		var reqs []TtsRequest
		_req := req
		var name string
		for k, v := range text {
			_req = req
			_req.Text = k
			_req.Params.Style = v
			name = name + k
			reqs = append(reqs, _req)
		}
		SSML := formatSSML1(reqs, i)
		post(SSML, i)
		//saveAzureSpeechFile(&_req, fmt.Sprintf("%s.mp3", name))
	}
}
func main() {
	req := [][]TtsRequest{
		{
			{
				Biz:    2,
				Source: 1,
				Text:   "听起来你很失望呢，",
				Lang:   "zh-CN",
				Params: TtsParams{
					Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
					Pitch:  "+5Hz",
					Rate:   "1.1",
					Volume: "90",
					Style:  "sad",
				},
			},
			{
				Biz:    2,
				Source: 1,
				Text:   "能告诉我发生了什么事吗？",
				Lang:   "zh-CN",
				Params: TtsParams{
					Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
					Pitch:  "+5Hz",
					Rate:   "1.1",
					Volume: "90",
					Style:  "cheerful",
				},
			},
		},
		//{
		//	{
		//		Biz:    2,
		//		Source: 1,
		//		Text:   "不行不行不行",
		//		Lang:   "zh-CN",
		//		Params: TtsParams{
		//			Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
		//			Pitch:  "+5Hz",
		//			Rate:   "1.1",
		//			Volume: "90",
		//			Style:  "chat",
		//		},
		//	},
		//	{
		//		Biz:    2,
		//		Source: 1,
		//		Text:   "为什么为什么为什么",
		//		Lang:   "zh-CN",
		//		Params: TtsParams{
		//			Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
		//			Pitch:  "+5Hz",
		//			Rate:   "1.1",
		//			Volume: "90",
		//			Style:  "cheerful",
		//		},
		//	},
		//},
		//{
		//	{
		//		Biz:    2,
		//		Source: 1,
		//		Text:   "不行不行不行",
		//		Lang:   "zh-CN",
		//		Params: TtsParams{
		//			Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
		//			Pitch:  "+5Hz",
		//			Rate:   "1.1",
		//			Volume: "90",
		//			Style:  "chat",
		//		},
		//	},
		//	{
		//		Biz:    2,
		//		Source: 1,
		//		Text:   "薬をのみます,txt:をのみます",
		//		Lang:   "zh-CN",
		//		Params: TtsParams{
		//			Voice:  "C002xiaosi_mor-stylesNeural_pro1Neural",
		//			Pitch:  "+5Hz",
		//			Rate:   "1.1",
		//			Volume: "90",
		//			Style:  "cheerful",
		//		},
		//	},
		//},
	}

	for i, text := range req {
		var name string
		var reqs []TtsRequest
		for j, v := range text {
			_req := req[i][j]
			_req.Text = v.Text
			_req.Params.Style = v.Params.Style
			name = name + v.Text
			reqs = append(reqs, _req)
		}
		SSML := formatSSML1(reqs, i)
		post(SSML, i)
		//saveAzureSpeechFile(&_req, fmt.Sprintf("%s.mp3", name))
	}
}

func post(xmlData string, i int) {
	xmlData = "<speak voice_type=\"xiaosi\" rate=\"1\" pitch=\"1\" volume=\"1\" emotion='confused'>\n    <break time=\"3000ms\"/>\n    屋漏偏逢连夜雨<break time=\"3000ms\"/>\n    船迟又遇打头风<break time=\"3000ms\"/>\n</speak>"
	payload := bytes.NewBufferString(xmlData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("X-Microsoft-OutputFormat", "audio-16khz-128kbitrate-mono-mp3")
	req.Header.Add("User-Agent", "curl")

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
	//流式返回
	//var (
	//	sendChan = make(chan TtsResponse, 10)
	//)
	//defer func() {
	//	fmt.Printf("close sendChan\n")
	//	close(sendChan)
	//}()
	//for{
	//	buf := make([]byte,1024)
	//	n,err := res.Body.Read(buf)
	//	if n > 0{
	//		sendChan <- TtsResponse{Result: buf[:n]}
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	// 将响应保存到文件中
	err = ioutil.WriteFile(fmt.Sprintf("output_%d.mp3", i), body, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	//var (
	//	sendChan = make(chan TtsResponse, 10)
	//)
	//defer func() {
	//	fmt.Printf("close sendChan\n")
	//	close(sendChan)
	//}()
	//sendChan <- TtsResponse{
	//	Result: body,
	//
	//}

	fmt.Println("Response saved to output.mp3")
}
func formatSSML1(reqs []TtsRequest, i int) string {
	var buf bytes.Buffer
	var d []Sentence
	for _, v := range reqs {
		data := Sentence{
			VoiceName:     v.Params.Voice,
			ExpressStyle:  v.Params.Style,
			ProsodyVolume: v.Params.Volume,
			ProsodyRate:   v.Params.Rate,
			ProsodyPitch:  v.Params.Pitch,
			Text:          v.Text,
		}
		d = append(d, data)
	}
	data := DataS{
		Sentences: d,
		Lang:      reqs[i].Lang,
		VoiceName: reqs[i].Params.Voice,
	}
	if err := ssmlFormat1.Execute(&buf, data); err != nil {
		ioutil.WriteFile(fmt.Sprintf("d.txt"), []byte("buf.Bytes()"), 0666)
		fmt.Printf("ssmlFormat.Execute failed: %v\n", err)
		return ""
	}
	ioutil.WriteFile(fmt.Sprintf("ssml_%d.txt", i), buf.Bytes(), 0666)
	fmt.Println(len(buf.String()))
	return buf.String()
}

var ssmlFormat1 = template.Must(template.New("azure").Parse(`<speak version="1.0" xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts" xml:lang="{{ .Lang }}">
	<voice name="{{ .VoiceName }}">			
		{{range .Sentences}}
		<mstts:express-as style="{{ .ExpressStyle }}">
			{{ if and .ProsodyVolume .ProsodyRate .ProsodyPitch }}
			<prosody volume="{{ .ProsodyVolume }}" rate="{{ .ProsodyRate }}" pitch="{{ .ProsodyPitch }}">{{ .Text }}</prosody>
			{{ else if and .ProsodyVolume .ProsodyRate }}
			<prosody volume="{{ .ProsodyVolume }}" rate="{{ .ProsodyRate }}">{{ .Text }}</prosody>
			{{ else if and .ProsodyVolume .ProsodyPitch }}
			<prosody volume="{{ .ProsodyVolume }}" pitch="{{ .ProsodyPitch }}">{{ .Text }}</prosody>
			{{ else if and .ProsodyRate .ProsodyPitch }}
			<prosody rate="{{ .ProsodyRate }}" pitch="{{ .ProsodyPitch }}">{{ .Text }}</prosody>
			{{ else if .ProsodyVolume }}
			<prosody volume="{{ .ProsodyVolume }}">{{ .Text }}</prosody>
			{{ else if .ProsodyRate }}
			<prosody rate="{{ .ProsodyRate }}">{{ .Text }}</prosody>
			{{ else if .ProsodyPitch }}
			<prosody pitch="{{ .ProsodyPitch }}">{{ .Text }}</prosody>
			{{ else }}
			{{ .Text }}
			{{ end }}
		</mstts:express-as>

		{{ end }}
    </voice>
</speak>`))
