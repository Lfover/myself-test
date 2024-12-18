package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func maskAllCharacters(input string) string {
	// 使用strings.Builder来构建新的字符串，以提高性能
	var result strings.Builder

	// 遍历输入字符串的每个字符
	for _, _ = range input {
		// 将每个字符替换为星号
		result.WriteRune('*')
	}

	// 返回替换后的字符串
	return result.String()
}

func main111() {
	// 测试方法
	reportData := "{\"learn_train_evaluation\":\"小明在学习上表现出了一定的积极性，特别是在查字词句方面，次数达到了4次，显示出对语言学习的兴趣。同时，他也进行了课程学习、背诵课文等多方面的尝试，表明他有广泛的学习兴趣。不过，建议小明可以增加背诵课文的频率来巩固记忆。\",\"advantage_content\":\"好奇;分析;学习\",\"advantage_analysis\":[\"- 好奇：询问了多个领域的知识，显示出强烈的求知欲。\",\"- 分析：通过查询不同类型的资料（如字词、诗歌等），展现了分析能力。\",\"- 学习：频繁查字词句体现了小明在语言领域持续学习的态度。\"],\"emotion_content\":\"冷静:50%;困惑:50%\",\"emotion_analysis\":\"孩子在对话中表现出对多个领域的好奇心，但在面对不熟悉的知识时也感到困惑。家长可以鼓励孩子保持好奇心，同时帮助其解决困惑，增强自信心。\",\"interest_content\":\"历史;文学\",\"interest_analysis\":\"孩子询问了二战的起因和杜甫的朝代，显示出对历史和文学的兴趣，可能是因为对这些话题感到好奇或想要了解更多。\",\"parent_advice\":[\"- 鼓励小明多阅读历史和文学作品，拓展知识面。\",\"- 支持小明在遇到困惑时主动寻求帮助，提高解决问题的能力。\"]}"
	reportData = strings.ReplaceAll(reportData, "\n", "")
	reportData = strings.ReplaceAll(reportData, "```json", "")
	reportData = strings.ReplaceAll(reportData, "```", "")
	input := "Hello"
	masked := maskAllCharacters(input)
	fmt.Println(masked) // 输出: **************
}
func main11() {
	for {
		source := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子
		ran := rand.New(source)                         // 生成一个rand
		fmt.Println(ran.Int())
		fmt.Println(ran.Int31())
		fmt.Println(ran.Intn(5))
		//const chars = "0123456789"
		//randomDigit := make([]byte, 20)
		//seed := time.Now().UnixNano()
		//
		//// 使用时间戳作为种子初始化随机数生成器
		//rand.Seed(seed)
		//for i := range randomDigit {
		//	source := rand.NewSource(time.Now().UnixNano()) // 使用当前的纳秒生成一个随机源，也就是随机种子
		//	ran := rand.New(source)
		//	randomDigit[i] = ran
		//}
		//fmt.Println(string(randomDigit))
		//time.Sleep(1 * time.Second)
	}

}
