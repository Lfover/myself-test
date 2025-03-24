package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := make(map[string]string)
	fmt.Println(len(s))
	n := len("26个英文字母是：A、B、C、D、E、F、G、H、T/J/KL")
	fmt.Println(n)
	tst := "我就是<speak version=\"1.0\" xmlns=\"http://www.w3.org/2001/10/synthesis\" xmlns:mstts=\"https://www.w3.org/2001/mstts\" xml:lang=\"en-US\">\n<voice name=\"en-GB-MaisieNeural\"><mstts:express-as style=\"cheerful\">this is a  \n<phoneme alphabet=\"ipa\" ph=\"nihao\">book</phoneme>\n \n</mstts:express-as></voice></speak>"
	a := ChunkSplitTxt1(tst, 2)
	for _, v := range a {
		fmt.Println("1111111111", v.Txt)
	}
	tst = "我就是<speak version=\"1.0\" xmlns=\"http://www.w3.org/2001/10/synthesis\" xmlns:mstts=\"https://www.w3.org/2001/mstts\" xml:lang=\"en-US\">\n<voice name=\"en-GB-MaisieNeural\"><mstts:express-as style=\"cheerful\">this is a  \n<phoneme alphabet=\"ipa\" ph=\"nihao\">book</phoneme>\n \n</mstts:express-as></voice></speak>，"
	a = ChunkSplitTxt1(tst, 2)
	for _, v := range a {
		s, _ := json.Marshal(v)
		fmt.Println("2222222222", string(s))
	}
}

type SplitTxt struct {
	Txt     string
	SsmlTxt string
	IsSsml  bool
}

func ChunkSplitTxt(txt string) []SplitTxt {
	var arr []SplitTxt

	if strings.Contains(txt, "<speak") {
		txtArr := SplitBySpeakTags(txt)
		for _, t := range txtArr {
			if strings.Contains(t, "<speak") {
				arr = append(arr, SplitTxt{
					Txt:     t,
					SsmlTxt: t,
					IsSsml:  true,
				})
			} else {
				arr = append(arr, SplitTxt{
					Txt:    t,
					IsSsml: false,
				})
			}
		}
	}

	return arr
}

func SplitBySpeakTags(text string) []string {
	re := regexp.MustCompile(`(?s)(<speak.*?</speak>)`)
	matches := re.FindAllStringIndex(text, -1)

	var result []string
	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			result = append(result, text[lastIndex:start])
		}

		result = append(result, text[start:end])
		lastIndex = end
	}

	if lastIndex < len(text) {
		result = append(result, text[lastIndex:])
	}

	return result
}
func ChunkSplitTxt1(txt string, action int32) []SplitTxt {
	var arr []SplitTxt
	if action == 3 {
		arr = []SplitTxt{{
			Txt:     txt,
			SsmlTxt: "",
			IsSsml:  false,
		},
		}
		return arr
	}
	if strings.Contains(txt, "<speak") {
		txtArr := SplitBySpeakTags(txt)
		for _, t := range txtArr {
			if strings.Contains(t, "<speak") {
				arr = append(arr, SplitTxt{
					Txt:     t,
					SsmlTxt: t,
					IsSsml:  true,
				})
			} else {
				chunkTxtArr := ChunkSplitTxt(t)
				for _, ct := range chunkTxtArr {
					arr = append(arr, SplitTxt{
						Txt:     ct.Txt,
						SsmlTxt: "",
						IsSsml:  false,
					})
				}
				arr = append(arr, SplitTxt{
					Txt:     t,
					SsmlTxt: "",
					IsSsml:  false,
				})
			}
		}
	} else {
		arr = []SplitTxt{{
			Txt:     txt,
			SsmlTxt: "",
			IsSsml:  false,
		},
		}
	}

	return arr
}
func getSplitSet(txt string) {
	//汉字

	return
}
