package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PromptResTal struct {
	ThinkingProcess string `json:"thinkingProcess"`
	Events          []struct {
		EventDate   string         `json:"eventDate"`
		EventType   string         `json:"eventType"`
		Person      []string       `json:"person"`
		Location    []string       `json:"location"`
		Title       string         `json:"title"`
		Description string         `json:"description"`
		ImpactScore interface{}    `json:"impactScore"`
		Emotion     map[string]int `json:"emotion"`
		//EmotionIntensity interface{}    `json:"emotionIntensity"`
		RecordDate string `json:"recordDate"`
		//RecordTime string `json:"record_time"`
	} `json:"events"`
	Interests map[string]int `json:"interestsList"`
	UserData  UserDataTal    `json:"userData"`
}
type UserDataTal struct {
	UserName string `json:"userName"`
}

func main() {
	s := "{'thinkingProcess': '用户多次询问二战的起因，表明对历史事件的兴趣。随后多次询问机器人的喜好，表明用户对机器人活动的兴趣。', 'events': [{'eventDate': '', 'eventType': '历史事件', 'person': [], 'location': [], 'title': '二战起因', 'description': '用户多次询问二战的起因，表现出对历史事件的强烈兴趣和好奇心。', 'impactScore': 3, 'emotion': {'好奇': 6}, 'recordDate': '2023-05-17'}, {'eventDate': '', 'eventType': '非事件', 'person': [], 'location': [], 'title': '询问机器人喜好', 'description': '用户多次询问机器人喜欢做什么，表现出对机器人活动的兴趣和好奇心。', 'impactScore': 2, 'emotion': {'好奇': 5}, 'recordDate': '2023-05-17'}], 'interestsList': {'历史事件': 6, '机器人活动': 5}, 'userData': {'userName': ''}}"
	a := strings.ReplaceAll(s, "'", "\"")
	chatRes := PromptResTal{}
	err := json.Unmarshal([]byte(a), &chatRes)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(chatRes)

}
