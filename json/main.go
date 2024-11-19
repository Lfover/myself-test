package main

import (
	"encoding/json"
	"fmt"
)

type Test1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type MarkResult struct {
	AsrResult              AsrResult   `bson:"asr_result" json:"asr_result" mapstructure:"asr_result"`
	AsrSemantic            AsrSemantic `bson:"asr_semantic" json:"asr_semantic" mapstructure:"asr_semantic"`
	HearXiaosi             bool        `bson:"hear_xiaosi" json:"hear_xiaosi" mapstructure:"hear_xiaosi"`
	InstructComplete       string      `bson:"instruct_complete" json:"instruct_complete" mapstructure:"instruct_complete"`
	IntentionRight         int         `json:"intention_right" bson:"intention_right" mapstructure:"intention_right"`
	Intention              Intention   `bson:"intention" json:"intention" mapstructure:"intention"`
	ResultType             bool        `bson:"result_type" json:"result_type" mapstructure:"result_type"`
	SecurityFiltering      string      `bson:"security_filtering" json:"security_filtering" mapstructure:"security_filtering"`
	Remark                 string      `bson:"remark" json:"remark" mapstructure:"remark"`
	SemanticClassification string      `bson:"semantic_classification" json:"semantic_classification" mapstructure:"semantic_classification"`
}
type AsrResult struct {
	Judgment int    `bson:"judgment" json:"judgment" mapstructure:"judgment"`
	Text     string `bson:"text" json:"text" mapstructure:"text"`
}
type AsrSemantic struct {
	Judgment int `bson:"judgment" json:"judgment" mapstructure:"judgment"`
}
type Intention struct {
	//Judgment       int    `bson:"judgment" json:"judgment" mapstructure:"judgment"`
	VerticalDomain string `bson:"vertical_domain" json:"vertical_domain" mapstructure:"vertical_domain"`
	Intention      string `bson:"intention" json:"intention" mapstructure:"intention"`
}

func main() {

	var aa MarkResult
	str := "{\\\"asr_result\\\":{\\\"judgment\\\":null,\\\"text\\\":\\\"\\\"},\\\"asr_semantic\\\":{\\\"judgment\\\":null},\\\"hear_xiaosi\\\":null}"
	json.Unmarshal([]byte(str), &aa)

	fmt.Println(aa.AsrResult.Judgment)
	fmt.Println(aa.AsrResult.Text)
	fmt.Println(aa.HearXiaosi)

}
