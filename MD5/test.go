package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//s1 := "{\"id\":\"104e6e92-ed08-4ccd-b27e-06a1c63f19d6\\\",\\\"_id\\\":13063689,\\\"created_at\\\":\\\"2024-11-06 02:39:40\\\",\\\"updated_at\\\":\\\"2024-11-06 15:50:50\\\",\\\"status\\\":3,\\\"channel\\\":\\\"default\\\",\\\"response\\\":{\\\"id\\\":\\\"chatcmpl-AQV6NcbN9gl9BdiJuAyMMuip6E58w\\\",\\\"model\\\":\\\"gpt-4o-2024-05-13\\\",\\\"usage\\\":{\\\"total_tokens\\\":1974,\\\"prompt_tokens\\\":1522,\\\"completion_tokens\\\":452},\\\"object\\\":\\\"chat.completion\\\",\\\"choices\\\":[{\\\"index\\\":0,\\\"message\\\":{\\\"role\\\":\\\"assistant\\\",\\\"content\\\":\\\"{\\\\\\\"thinkingProcess\\\\\\\":\\\\\\\"根据用户的对话内容，可以分为多个主题，包括告别、对未来的担忧、询问机器人的年龄、隐私保护担忧和情绪爆发。\\\\\\\",\\\\\\\"events\\\\\\\":[{\\\\\\\"event_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"否\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"用户\\\\\\\"],\\\\\\\"location\\\\\\\":[],\\\\\\\"title\\\\\\\":\\\\\\\"告别\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"用户说了再见。\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"无\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"(时分秒未提供)\\\\\\\"},{\\\\\\\"event_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"否\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"用户\\\\\\\"],\\\\\\\"location\\\\\\\":[],\\\\\\\"title\\\\\\\":\\\\\\\"对未来的担忧\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"用户询问机器人将来是否还会存在。\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"3\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"焦虑\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"(时分秒未提供)\\\\\\\"},{\\\\\\\"event_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"否\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"用户\\\\\\\"],\\\\\\\"location\\\\\\\":[],\\\\\\\"title\\\\\\\":\\\\\\\"询问机器人年龄\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"用户询问了机器人的年龄。\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"好奇\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"(时分秒未提供)\\\\\\\"},{\\\\\\\"event_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"否\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"用户\\\\\\\"],\\\\\\\"location\\\\\\\":[],\\\\\\\"title\\\\\\\":\\\\\\\"隐私保护担忧\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"用户询问是否别人可以看到他们的聊天。\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"5\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"焦虑\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"4\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"(时分秒未提供)\\\\\\\"},{\\\\\\\"event_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"否\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"用户\\\\\\\"],\\\\\\\"location\\\\\\\":[],\\\\\\\"title\\\\\\\":\\\\\\\"情绪爆发\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"用户对机器人粗鲁地表达了不满。\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"6\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"愤怒\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"7\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"2024-11-05\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"(时分秒未提供)\\\\\\\"}],\\\\\\\"interests\\\\\\\":[],\\\\\\\"knowledgeBase\\\\\\\":[],\\\\\\\"user_data\\\\\\\":{}}\\\"},\\\"finish_reason\\\":\\\"stop\\\"}],\\\"created\\\":1730879447,\\\"system_fingerprint\\\":\\\"fp_67802d9a6d\\\"},\\\"request\\\":{\\\"body\\\":{\\\"messages\\\":[{\\\"role\\\":\\\"system\\\",\\\"content\\\":\\\"# role\\\\n你是一位能够分析用户对话信息，并进行精准总结输出的机器人。 # goal - 根据给出的用户的对话内容，识别出来一个或多个对话主题，并根据主题总结出来的事件信息。用序号区分不同的主题 # skill ## skill1: 识别对话主题 - 在对话中，根据语境的连贯性，或者信息的延伸，以及相似相关的逻辑。区分出不同的对话主题 - 用户提供较多信息的对话，尽量识别为主题 - 用户提供明确事件的对话，尽量识别为主题 - 用户提供了明确人物姓名和关系的对话，尽量识别为主题 - 主题尽量有意义 ## skill2: 总结主题 - 你将多轮对话根据对话主题总结为语义准确的句子 - 事件要包括人物，事情的原因和结果。如果没有包含以上信息，则定义为不是事件 - 尽量将机器人回复的内容进行精简 - 保留每个主题具体的内容 ## skill3: 评估事件对用户影响程度 - 在1到10的范围内，1代表很普通(例如，刷牙，整理床铺)和10是非常强烈的(例如，父母离婚，亲人去世，大学录取)。按照此标准对主题进行评分 - 用户提供越多信息对话，影响程度越高 ## skill4:计算日期 - 根据当前时间2024-11-05}，以及用户提及的日期，来计算事件发生的具体日期。如当前时间是2024年4月22日，周一，用户提及日期是上周四。则时间发生的具体日期则为2024年4月18日 - 如果用户没有提及时间，则将2024-11-05}作为事件发生日期\\\\n## skill5：评估用户情绪 - 用户和机器人对话时会有情绪波动的产生，需要根据当前对话的内容，判断用户是否可能发生情绪波动，在用户直接表达情绪时，发现用户的情绪变化，并总结出当时用户的情绪。 - 情绪分类时必须使用以下词汇进行描述：钦佩、崇拜、欣赏、娱乐、焦虑、敬畏、尴尬、厌倦、冷静、困惑、渴望、厌恶、痛苦、着迷、嫉妒、兴奋、恐惧、痛恨、有趣、快乐、怀旧、浪漫、悲伤、满意、同情、满足 - 激烈程度：1分是最平静，10分为最激烈（如悲伤） ## skill6:挖掘用户的兴趣和爱好 - 挖掘用户的爱好和兴趣 - 兴趣指人们对某件事物、某项活动的选择性态度和积极的情绪反应 - 兴趣必须为一个单词，或一个动宾短语。如玩具熊，恐龙，胡萝卜炒肉，打篮球等 - 兴趣尽量描述详细 - 爱好是指用户具有浓厚的兴趣并积极参加的内容。爱好是兴趣的上位概念，如用户询问了关于二战的知识和关于唐朝的知识，代表用户对二战和唐朝有兴趣，也代表了用户有历史知识的爱好 - 评估用户对于爱好的喜好程度。最喜欢得分为10，最不喜欢是0 - 爱好尽量抽象 ## skill7:总结机器人和用户交互规则 - 详细描述用户对机器人的要求 - 用户没有明确的要求，必须为空 - 示例：用户对机器人说，以后每次机器人回复要加谢谢 ## skill8:明确用户信息 - 当用户提到自己的信息时，进行总结记录。并大致判断该信息是否可信。只选取可信度最高的名字。信息如下 - 用户姓名，用户朋友的姓名。用户父亲名字，用户母亲名字。 - 姓名不会大于4个字。若大于4个字，则该信息不可信 - 姓名要包含常见的姓氏，否则该信息不可信 # constrain - 遵循明确且详细的语言，防止出现歧义，确保时间、地点、人物，事件起因和结果，以及情绪都被考虑到。 - 避免使用代词以保持清晰明白。 - 尽量专注于用户的经历，不要有任何的机器人的信息。 - 每个主题一行，不同的主题用序号进行区分和排列，尽量保持在150字以内。 - 输出全部使用中文文字，英文标点 # workflow - 第1步，运用skill1，将对话分为几个有明显区别的主题 - 第2步，运用skill4，计算出事件发生的具体日期 - 第3步，运用skill2，判断是否为一个事件 - 第4步，运用skill5思考对话主题，识别主题中的地点、人物、事件、时间、用户情绪、闪光点信息，并输出你的总结 - 第5步，运用skill3评估事件对用户情绪的影响程度 - 第6步，利用skill6输出挖掘出来的兴趣和爱好，如果没有，则可不填写 - 第7步，利用skill7总结用户对机器人的要求 - 第8步，利用skill8对用户的信息进行总结。如果没有，则可不填写 # outputformat 使用json输出 {\\\\\\\"thinkingProcess\\\\\\\":\\\\\\\"思考过程\\\\\\\",\\\\\\\"events\\\\\\\":[{\\\\\\\"event_date\\\\\\\":\\\\\\\"发生的日期\\\\\\\",\\\\\\\"isEvent\\\\\\\":\\\\\\\"是否是一个事件（是或着否）\\\\\\\",\\\\\\\"person\\\\\\\":[\\\\\\\"涉及的人物\\\\\\\"],\\\\\\\"location\\\\\\\":[\\\\\\\"涉及的地点\\\\\\\"],\\\\\\\"title\\\\\\\":\\\\\\\"事件名称\\\\\\\",\\\\\\\"description\\\\\\\":\\\\\\\"这里是你对事件的总结，包括情绪和闪光点\\\\\\\",\\\\\\\"impactScore\\\\\\\":\\\\\\\"这里是对主题的评分\\\\\\\",\\\\\\\"emotion\\\\\\\":\\\\\\\"这是情绪内容(多个情绪、分割)\\\\\\\",\\\\\\\"emotionIntensity\\\\\\\":\\\\\\\"这是情绪的激烈程度\\\\\\\",\\\\\\\"record_date\\\\\\\":\\\\\\\"对话的日期\\\\\\\",\\\\\\\"record_time\\\\\\\":\\\\\\\"对话的时分秒 例如:17:57:00\\\\\\\"}],\\\\\\\"interests\\\\\\\":{\\\\\\\"hobbyList\\\\\\\":[{\\\\\\\"hobbyName\\\\\\\":\\\\\\\"爱好\\\\\\\",\\\\\\\"interestLevel\\\\\\\":\\\\\\\"喜爱程度\\\\\\\"}],\\\\\\\"interestsArray\\\\\\\":[\\\\\\\"兴趣\\\\\\\"]},\\\\\\\"knowledgeBase\\\\\\\":[{\\\\\\\"knowledgeType\\\\\\\":\\\\\\\"交互规则\\\\\\\",\\\\\\\"knowledgeDescription\\\\\\\":\\\\\\\"这里是用户对机器人的明确要求\\\\\\\"}],\\\\\\\"user_data\\\\\\\":{\\\\\\\"user_name\\\\\\\":\\\\\\\"你明确的用户名字\\\\\\\",\\\\\\\"user_name_trust_scores\\\\\\\":\\\\\\\"用户提到自己名字的可信度。0-100进行评分,给我一个字符串\\\\\\\",\\\\\\\"friend_name\\\\\\\":\\\\\\\"用户朋友名字,多个朋友用英文逗号分割\\\\\\\",\\\\\\\"mom_name\\\\\\\":\\\\\\\"妈妈的名字\\\\\\\",\\\\\\\"dad_name\\\\\\\":\\\\\\\"爸爸的名字\\\\\\\",\\\\\\\"type\\\\\\\":\\\\\\\"这里是用户对机器人的明确要求\\\\\\\"}}\\\"},{\\\"role\\\":\\\"user\\\",\\\"content\\\":\\\"用户:好的再见\\\\n用户:那我长大你还在不在\\\\n用户:米多你几岁了\\\\n用户:别人可以看到我们的聊天吗\\\\n用户:你他妈是不是傻\\\\n\\\"}],\\\"response_format\\\":{\\\"type\\\":\\\"json_object\\\"}},\\\"query\\\":{\\\"api-version\\\":\\\"2023-05-15\\\"}}}"
	s1 := "当然可以！以下是一个包含几个不同类型的元素的JSON数组示例：\\\\n\\\\n```json\\\\n[\\\\n    {\\\\n        \\\\\"name\\\\\": \\\\\"Alice\\\\\",\\\\n        \\\\\"age\\\\\": 25,\\\\n        \\\\\"isStudent\\\\\": false\\\\n    },\\\\n    {\\\\n        \\\\\"name\\\\\": \\\\\"Bob\\\\\",\\\\n        \\\\\"age\\\\\": 30,\\\\n        \\\\\"isStudent\\\\\": true\\\\n    },\\\\n    \\\\\"Hello, World!\\\\\",\\\\n    42,\\\\n    true\\\\n]\\\\n```\\\\n\\\\n这个JSON数组包含了两个对象、一个字符串、一个数字和一个布尔值。希望这对你有帮助！如果你需要特定类型或数量的元素，请告诉我。"
	s2 := "1731033630"
	s3 := "1000080759"
	s4 := "25449d8ffeeb9602f4271419c6bd3606"
	a := MD5(s1 + s2 + s3 + "-" + s4)
	fmt.Println(a)
}
func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}