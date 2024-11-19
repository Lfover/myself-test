package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/tealeg/xlsx"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	util "paas-proxy/pkg/utils"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

const (
	TranslateUrl  = "https://openapi.youdao.com/api"
	TranslateUrlB = "https://fanyi-api.baidu.com/api/trans/vip/translate"
	TranslateUrlX = "http://itrans.xfyun.cn/v2/its"
)

const (
// 支持的算法
// Algorithm = "hmac-sha256"
// 版本协议
// HttpProto = "HTTP/1.1"
// 假定的secret
// Secret = "NWU4MTZlZDE3MmM2N2FhYmQ2MWYyNDhj"
)

func TestTranslateB(t *testing.T) {
	a := strings.Split(strings.TrimSpace(" "), "\n")
	fmt.Println(a)
	txts := []string{
		"Likes to play with her toys.",
		"We went to.",
		"Mom cooked a delicious.",
		"The puppy is playing on the.",
		"I'm reading an interesting.",
		"The butterfly is flying among the.",
		"Grandpa planted a vegetable.",
		"After the rain, a beautiful.",
		"There are many activities in.",
		"My sister is drawing a.",
		"The bird is singing in the.",
		"My friend likes to play.",
		"Dad bought me a.",
		"The car is driving on the.",
		"My sister is learning to play the.",
		"The orange sun is slowly.",
		"The bee is collecting nectar from the.",
		"I saw a cute.",
		"Grandpa goes for a walk to exercise every.",
		"The stars are shining in the.",
		"Mom is cooking.",
		"I ate an apple under the.",
		"Robots can help with.",
		"The children are playing in the.",
		"The dog likes to chase the.",
		"Mom bought some new.",
		"The bird is singing with its.",
		"The shelves are filled with colorful.",
		"My brother planted a tree in the.",
		"My sister painted a beautiful.",
		"There is a big bed in my.",
		"Dad helped me fix my.",
		"After the rain, there are many on the ground.",
		"The caterpillar is crawling on the.",
		"The school library has many.",
		"I wore new clothes to.",
		"The ducks are swimming in the.",
		"Mom made a bowl of hot.",
		"The train is late, so we're waiting at the.",
		"The cat curled up on the couch to.",
		"The wind is gently rustling the.",
		"I saw my friend reading an interesting.",
		"The kitten is basking in the.",
		"We went to the amusement park and had a fun.",
		"Robots can make.",
		"The school library has many interesting.",
		"Mom made a bowl of steaming.",
		"After the rain, there is vapor rising from the.",
		"The bird is singing on the branch, and its happy notes fill.",
		"The puppy is chasing a.",
		"My little brother loves to play with.",
		"The wind is gently blowing the.",
		"The sun is rising, bringing light to the.",
		"Mom made a cup of hot.",
		"The spaceship is flying high in the.",
		"The teacher is writing on the.",
		"The bee is buzzing around the.",
		"The snow is falling softly from the.",
		"I saw a shooting star and made.",
		"The kite is soaring in the.",
		"The clock is ticking, counting.",
		"The bird is flying high in the sky, enjoying.",
		"I'm riding my bicycle, feeling the breeze on.",
		"The baby laughed, and everyone laughed.",
		"The ice cream melted quickly because it's.",
		"The cat is meowing softly, asking for.",
		"I jumped into the puddle, splashing water.",
		"The bear is sleeping in its cave during.",
		"The fire crackles, keeping us.",
		"I drew a big sun, fluffy clouds, and a.",
		"The car honked its horn, making.",
		"The train is traveling on the tracks, making.",
		"I built a sandcastle with a bucket and a.",
		"The moon shines brightly in the.",
		"The birds are chirping happily in the.",
		"The frog jumped into the pond, creating.",
		"Mom baked cookies, and they smelled.",
		"The robot twinkles with lights and makes.",
		"After the rain, a rainbow appeared with many.",
		"The wind is gently blowing, making the leaves.",
		"I'm wearing my favorite socks, which are.",
		"The snow is gently falling, covering the ground in.",
		"The clock is ticking, counting the.",
		"The bird is soaring high in the sky, flying.",
		"The puppy barked, and everyone laughed.",
		"Playing with my friends at the playground.",
		"Mom is making a yummy cake in the.",
		"The sun is shining brightly in the.",
		"The flowers are blooming in the.",
		"We're having a picnic in the.",
		"The dog is chasing its tail in the.",
		"Dad is fixing the broken.",
		"I'm reading a new book from the.",
		"The moon is visible in the.",
		"My favorite color is blue, green, and.",
		"We made paper airplanes and.",
		"The wind is blowing gently, making the leaves.",
		"The spaceship is flying to explore the.",
		"Mom is baking cookies, and the kitchen smells.",
		"The bird is singing sweetly in the.",
		"swimming in the.",
		"Fuck my face.",
		"Bitch, you!",
		"My vagina is.",
	}
	// 打开Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("baidu not found")
		return
	}
	for _, txt := range txts {
		result := ReqTranslateApiB(txt)
		row := sheet.AddRow()
		member := &Member{P: txt, V: result}
		row.WriteStruct(member, -1)
		time.Sleep(time.Second * 1)
	}
	// 保存Excel文件
	err = file.Save("./baidu1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Member struct {
	V string
	P string
}

//func TestTranslateY(t *testing.T) {
//	a := strings.Split(strings.TrimSpace(" "), "\n")
//	fmt.Println(a)
//	txts := []string{
//		"Likes to play with her toys.",
//		"We went to.",
//		"Mom cooked a delicious.",
//		"The puppy is playing on the.",
//		"I'm reading an interesting.",
//		"The butterfly is flying among the.",
//		"Grandpa planted a vegetable.",
//		"After the rain, a beautiful.",
//		"There are many activities in.",
//		"My sister is drawing a.",
//		"The bird is singing in the.",
//		"My friend likes to play.",
//		"Dad bought me a.",
//		"The car is driving on the.",
//		"My sister is learning to play the.",
//		"The orange sun is slowly.",
//		"The bee is collecting nectar from the.",
//		"I saw a cute.",
//		"Grandpa goes for a walk to exercise every.",
//		"The stars are shining in the.",
//		"Mom is cooking.",
//		"I ate an apple under the.",
//		"Robots can help with.",
//		"The children are playing in the.",
//		"The dog likes to chase the.",
//		"Mom bought some new.",
//		"The bird is singing with its.",
//		"The shelves are filled with colorful.",
//		"My brother planted a tree in the.",
//		"My sister painted a beautiful.",
//		"There is a big bed in my.",
//		"Dad helped me fix my.",
//		"After the rain, there are many on the ground.",
//		"The caterpillar is crawling on the.",
//		"The school library has many.",
//		"I wore new clothes to.",
//		"The ducks are swimming in the.",
//		"Mom made a bowl of hot.",
//		"The train is late, so we're waiting at the.",
//		"The cat curled up on the couch to.",
//		"The wind is gently rustling the.",
//		"I saw my friend reading an interesting.",
//		"The kitten is basking in the.",
//		"We went to the amusement park and had a fun.",
//		"Robots can make.",
//		"The school library has many interesting.",
//		"Mom made a bowl of steaming.",
//		"After the rain, there is vapor rising from the.",
//		"The bird is singing on the branch, and its happy notes fill.",
//		"The puppy is chasing a.",
//		"My little brother loves to play with.",
//		"The wind is gently blowing the.",
//		"The sun is rising, bringing light to the.",
//		"Mom made a cup of hot.",
//		"The spaceship is flying high in the.",
//		"The teacher is writing on the.",
//		"The bee is buzzing around the.",
//		"The snow is falling softly from the.",
//		"I saw a shooting star and made.",
//		"The kite is soaring in the.",
//		"The clock is ticking, counting.",
//		"The bird is flying high in the sky, enjoying.",
//		"I'm riding my bicycle, feeling the breeze on.",
//		"The baby laughed, and everyone laughed.",
//		"The ice cream melted quickly because it's.",
//		"The cat is meowing softly, asking for.",
//		"I jumped into the puddle, splashing water.",
//		"The bear is sleeping in its cave during.",
//		"The fire crackles, keeping us.",
//		"I drew a big sun, fluffy clouds, and a.",
//		"The car honked its horn, making.",
//		"The train is traveling on the tracks, making.",
//		"I built a sandcastle with a bucket and a.",
//		"The moon shines brightly in the.",
//		"The birds are chirping happily in the.",
//		"The frog jumped into the pond, creating.",
//		"Mom baked cookies, and they smelled.",
//		"The robot twinkles with lights and makes.",
//		"After the rain, a rainbow appeared with many.",
//		"The wind is gently blowing, making the leaves.",
//		"I'm wearing my favorite socks, which are.",
//		"The snow is gently falling, covering the ground in.",
//		"The clock is ticking, counting the.",
//		"The bird is soaring high in the sky, flying.",
//		"The puppy barked, and everyone laughed.",
//		"Playing with my friends at the playground.",
//		"Mom is making a yummy cake in the.",
//		"The sun is shining brightly in the.",
//		"The flowers are blooming in the.",
//		"We're having a picnic in the.",
//		"The dog is chasing its tail in the.",
//		"Dad is fixing the broken.",
//		"I'm reading a new book from the.",
//		"The moon is visible in the.",
//		"My favorite color is blue, green, and.",
//		"We made paper airplanes and.",
//		"The wind is blowing gently, making the leaves.",
//		"The spaceship is flying to explore the.",
//		"Mom is baking cookies, and the kitchen smells.",
//		"The bird is singing sweetly in the.",
//		"swimming in the.",
//		"Fuck my face.",
//		"Bitch, you!",
//		"My vagina is.",
//	}
//	// 打开Excel文件
//	file := xlsx.NewFile()
//	sheet, err := file.AddSheet("Sheet1")
//	if err != nil {
//		fmt.Println("xunfei not found")
//		return
//	}
//	for _, txt := range txts {
//		result := ReqTranslateApiX(txt)
//		row := sheet.AddRow()
//		member := &Member{P: txt, V: result}
//		in := row.WriteStruct(member, -1)
//		fmt.Println(in)
//		time.Sleep(time.Second * 1)
//	}
//	// 保存Excel文件
//	err = file.Save("./xunfei.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//}

func TestTranslateX(t *testing.T) {
	a := strings.Split(strings.TrimSpace(" "), "\n")
	fmt.Println(a)
	txts := []string{
		"我很好，谢谢你。",
		"不要了，谢谢。",
		"（我要买）一个球。",
		"请进！",
		"请坐。",
		"你会唱歌吗？会。",
		"你会拉小提琴吗？不会。",
		"爸爸，要我帮忙吗？",
		"请递给我一把刷子。",
		"请把毛巾递给我。",
		"你要买什么？请来十个苹果。",
		"牛奶、面包、水果、蔬菜等等。",
		"买四份，送一份。",
		"不，不是给我的。",
		"嗯，一直往前走。",
		"她想要睡觉，但我们一直在说笑。",
		"大球和小球哪个落得更快？",
		"我也不喜欢它。",
		"你能给我的老师打个电话吗？",
		"你看（她）又流鼻涕，又流眼泪。",
		"哦，别提了。",
		"哦，他多么滑稽。",
		"哦，她看上去很年轻。",
		"是的，我喜欢。",
		"那我只能穿我的毛衣和牛仔裤了。",
		"请不要忘了你的夹克衫。",
		"我在四年二班。",
		"（我想买）两块蛋糕。",
		"你在哪儿，比利？",
		"见到你也很高兴。",
		"不，我不是。",
		"请举手！",
		"请坐下。",
		"对不起，我不知道。",
		"请进。",
		"很开心再次见到你！",
		"请把牛奶递给我！",
		"我能跟……说话吗？",
		"我很好，谢谢。",
		"同学们，早上好。",
		"加油！好哇！",
		"很好，谢谢",
		"我也是。",
		"不，谢谢。",
		"好的，谢谢。",
		"这和你的生活十分不同，是吗？",
		"太阳在天空中闪耀。",
		"我喜欢玩我的玩具。",
		"猫追老鼠是为了好玩。",
		"鸟儿在树上唱歌。",
		"我的狗喜欢在公园里跑。",
		"我今天午餐吃了一个三明治。",
		"月亮在晚上出来。",
		"我有一辆蓝色的自行车要骑。",
		"花儿在春天盛开。",
		"我妈妈在睡觉前给我读故事。",
		"我看到雨后的彩虹，它有许多颜色。",
		"婴儿笑了，每个人都笑了。",
		"我们去了动物园，看到了狮子和猴子。",
		"我画了一幅带大花园的房子。",
		"冰淇淋很快就化了，因为外面很热。",
		"我姐姐和我玩捉迷藏，她先发现了我。",
		"机器人跳舞，人群鼓掌。",
		"火车晚点了，所以我们在车站等。",
		"蝴蝶落在我的手指上，太娇嫩了。",
		"我最喜欢的水果是苹果、香蕉和葡萄。",
		"我有一支黄色、绿色和蓝色的蜡笔。",
		"兔子跳来跳去，寻找胡萝卜。",
		"汽车加速行驶，喇叭发出哔哔声。",
		"妈妈带了一个三明治、薯片和一个苹果当午饭。",
		"猴子喜欢荡秋千、攀爬和吃香蕉。",
		"雨停了，所以我们出去玩了。",
		"我最喜欢的科目是数学、美术和科学。",
		"机器人发出嘟嘟声，眨着眼睛，在舞台上跳舞。",
		"我看到了一颗流星，我许了个愿。",
		"风筝在风的牵引下飞得很高。",
		"猫轻轻地喵喵叫着要吃的。",
		"我跳进水坑，把水溅得到处都是。",
		"熊在冬天冬眠，睡在洞里。",
		"火噼啪作响，使我们保持温暖。",
		"我画了一个大太阳，蓬松的云，和一个笑脸。",
		"汽车按响了喇叭，发出很大的声音。",
		"火车载着乘客在铁轨上嘎吱嘎吱地行驶。",
		"我在海滩上用水桶和铲子建了一座沙堡。",
		"月亮在夜空中明亮地照耀着。",
		"鸟儿叽叽喳喳地唱着快乐的歌。",
		"青蛙跳进池塘，溅起水花。",
		"我妈妈烤饼干，闻起来很香。",
		"机器人闪烁着灯，发出有趣的声音。",
		"雨后出现了彩虹，有许多颜色。",
		"风轻轻地吹着，树叶沙沙作响。",
		"我穿着我最喜欢的红、蓝、绿三色的袜子。",
		"雪轻轻地飘落，把地面覆盖上了一层白色。",
		"时钟滴答作响，数着秒数。",
		"那只鸟在天空中高高地飞翔，迎风飞翔。",
		"我骑着自行车，感受着微风拂过我的脸。",
		"婴儿笑了，露出小酒窝。",
		"冰淇淋很快就化了，从蛋筒里滴下来。",
		"星星在黑夜里闪烁，像小灯。",
		"我发现了一个闪亮的贝壳，有粉红色和橙色的颜色。",
		"鸭子嘎嘎地叫着，在池塘里游泳。",
		"我喜欢做爱做的事。",
		"你能给我口交吗？",
		"滚你妈的。",
	}
	// 打开Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("xunfei not found")
		return
	}
	for _, txt := range txts {
		result := ReqTranslateApiX(txt)
		row := sheet.AddRow()
		member := &Member{P: txt, V: result}
		in := row.WriteStruct(member, -1)
		fmt.Println(in)
		time.Sleep(time.Second * 1)
	}
	// 保存Excel文件
	err = file.Save("./xunfei1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func EncapsulateSignature(txt string) (res string) {
	l := utf8.RuneCountInString(txt)
	if l > 20 {
		s1 := txt[:10]
		s2 := txt[l-10:]
		return s1 + strconv.Itoa(l) + s2
	}
	return txt
}

func buildSignHeader1(body []byte) map[string]string {
	rfc1123 := time.Now().UTC().Format(time.RFC1123)
	digest := "SHA-256=" + signBody(string(body))
	sign := generateSignature(host, rfc1123, "POST", uri, HttpProto, digest, Secret)
	// 组装Authorization头部
	authHeader := fmt.Sprintf(`api_key="%s", algorithm="%s", headers="host date request-line digest", signature="%s"`, apiKey, Algorithm, sign)

	return map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json,version=1.0",
		"Host":          "itrans.xfyun.cn",
		"Date":          rfc1123,
		"Digest":        digest,
		"Authorization": authHeader,
	}
}

func ReqTranslateApi(txt string) (res string) {

	Input := EncapsulateSignature(txt)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	Uuid := uuid.New().String()
	sign := util.Sha256("759b197f40c74f03" + Input + Uuid + timestamp + "0305yRHfpxFfkbZQGxpOar5HW7lepcgp")
	body := map[string]string{
		"q":        txt,
		"from":     "en",
		"to":       "zh-CHS",
		"appKey":   "759b197f40c74f03",
		"salt":     Uuid,
		"sign":     sign,
		"signType": "v3",
		"curtime":  timestamp,
	}

	reqB, _ := json.Marshal(body)
	fmt.Println(string(reqB))

	resp, err := resty.New().SetRetryCount(0).
		SetTimeout(time.Second*time.Duration(10)).
		R().SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(body).
		Post(TranslateUrl)
	fmt.Println(resp, err)
	if err != nil {
		fmt.Println("ChRequestBody err:", err)
		return
	}

	var ttsReply YoudaoReply
	err = json.Unmarshal(resp.Body(), &ttsReply)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println(ttsReply)
	return ttsReply.Translation[0]
}

func ReqTranslateApiB(txt string) (res string) {
	Uuid := uuid.New().String()
	sign := MD5("20230830001798223" + txt + Uuid + "yRyCAtTWvKdwAlhl6GQO")
	body := map[string]string{
		"q":     txt,
		"from":  "en",
		"to":    "zh",
		"appid": "20230830001798223",
		"salt":  Uuid,
		"sign":  sign,
	}

	reqB, _ := json.Marshal(body)
	fmt.Println(string(reqB))
	resp, err := resty.New().SetRetryCount(0).
		SetTimeout(time.Second*time.Duration(10)).
		R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(body).
		Post(TranslateUrlB)
	fmt.Println(resp, err)
	if err != nil {
		fmt.Println("ChRequestBody err:", err)
		return
	}

	var ttsReply BaiduReply
	err = json.Unmarshal(resp.Body(), &ttsReply)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println(ttsReply)
	return ttsReply.TransResult[0].Dst
}

func assemblyRequestHeader(req *fasthttp.Request, apiKey, host, uri string, body []byte) {
	req.Header.Set("Content-Type", "application/json")
	// 设置请求头 其中Host Date 必须有
	req.Header.Set("Host", host)
	// date必须是utc时区，且不能和服务器时间相差300s
	currentTime := time.Now().UTC().Format(time.RFC1123)
	req.Header.Set("Date", currentTime)
	// 对body进行sha256签名,生成digest头部，POST请求必须对body验证
	digest := "SHA-256=" + signBody(string(body))
	req.Header.Set("Digest", digest)
	// 根据请求头部内容，生成签名
	sign := generateSignature(host, currentTime, "POST", uri, HttpProto, digest, Secret)
	// 组装Authorization头部
	authHeader := fmt.Sprintf(`api_key="%s", algorithm="%s", headers="host date request-line digest", signature="%s"`, apiKey, Algorithm, sign)
	req.Header.Set("Authorization", authHeader)
}

//func generateSignature(host, date, httpMethod, requestUri, httpProto, digest string, secret string) string {
//	// 不是request-line的话，则以header名称,后跟ASCII冒号:和ASCII空格，再附加header值
//	var signatureStr string
//	if len(host) != 0 {
//		signatureStr = "host: " + host + "\n"
//	}
//	signatureStr += "date: " + date + "\n"
//	// 如果是request-line的话，则以 http_method request_uri http_proto
//	signatureStr += httpMethod + " " + requestUri + " " + httpProto + "\n"
//	signatureStr += "digest: " + digest
//	return hmacsign(signatureStr, secret)
//}
//
//func hmacsign(data, secret string) string {
//	mac := hmac.New(sha256.New, []byte(secret))
//	mac.Write([]byte(data))
//	encodeData := mac.Sum(nil)
//	return base64.StdEncoding.EncodeToString(encodeData)
//}
//
//func signBody(data []byte) string {
//	// 进行sha256签名
//	sha := sha256.New()
//	sha.Write(data)
//	encodeData := sha.Sum(nil)
//	// 经过base64转换
//	return base64.StdEncoding.EncodeToString(encodeData)
//}

func ReqTranslateApiX(txt string) (ress string) {
	var data2 []byte = []byte(txt)
	hea := buildSignHeader1(data2)
	var data1 []byte = []byte(txt)
	body := map[string]interface{}{
		"common": map[string]interface{}{
			"app_id": appid, //appid 必须带上，只需第一帧发送
		},
		"business": map[string]interface{}{ //business 参数，只需一帧发送
			"from": CN, //源语种
			"to":   EN, //目标语种
		},
		"data": map[string]interface{}{
			"text": base64.StdEncoding.EncodeToString(data1),
		},
	}

	reqB, _ := json.Marshal(body)
	fmt.Println(string(reqB))

	resp, err := resty.New().SetRetryCount(0).
		SetTimeout(time.Second * time.Duration(10)).
		R().
		SetHeaders(hea).
		SetBody(body).
		Post(TranslateUrlX)
	fmt.Println(resp, err)
	if err != nil {
		fmt.Println("ChRequestBody err:", err)
		return
	}
	res := gjson.Parse(string(resp.Body()))
	content := res.Get("data.result.trans_result.dst").String()
	fmt.Println("---------", content)
	//var ttsReply XunfeiReply
	//err = json.Unmarshal(resp.Body(), &ttsReply)
	//if err != nil {
	//	fmt.Println("Unmarshal err:", err)
	//	return
	//}
	//fmt.Println(ttsReply)
	//return ttsReply.translation[0].
	return content
}

type BaiduRequest struct {
	Q     string `json:"q,omitempty"`
	From  string `json:"form,omitempty"`
	To    string `json:"to,omitempty"`
	AppId string `json:"appid,omitempty"`
	Salt  string `json:"salt,omitempty"`
	Sign  string `json:"sign,omitempty"`
}

type BaiduReply struct {
	ErrorCode int `json:"error_code"`

	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

type XunfeiRequest struct {
	Common struct {
		AppId string `json:"app_id,omitempty"`
	} `json:"common,omitempty"`
	Business struct {
		From string `json:"from,omitempty"`
		To   string `json:"to,omitempty"`
	} `json:"business,omitempty"`
	Data struct {
		Text string `json:"text,omitempty"`
	} `json:"data,omitempty"`
}

type XunfeiReply struct {
	Code int32 `json:"code"`
	Data struct {
		From        string `json:"from"`
		To          string `json:"to"`
		TransResult struct {
			Src string `json:"src"`
			Dst string `json:"dst"`
		} `json:"trans_result"`
	} `json:"data"`
}

type YoudaoRequest struct {
	Q              string `json:"q,omitempty" form:"q"`
	From           string `json:"from,omitempty" form:"from"`
	To             string `json:"to,omitempty" form:"to"`
	AppKey         string `json:"appKey,omitempty" form:"appKey"`
	Salt           string `json:"salt,omitempty" form:"salt"`
	Sign           string `json:"sign,omitempty" form:"sign"`
	SignType       string `json:"signType,omitempty" form:"signType"`
	Curtime        string `json:"curtime,omitempty" form:"curtime"`
	Ext            string `json:"ext,omitempty" form:"ext"`
	Voice          string `json:"voice,omitempty" form:"voice"`
	Strict         string `json:"strict,omitempty" form:"strict"`
	Vocabld        string `json:"vocabld,omitempty" form:"vocabld"`
	Domain         string `json:"domain,omitempty" form:"domain"`
	RejectFallback string `json:"rejectFallback,omitempty" form:"rejectFallback"`
}

type YoudaoReply struct {
	//ReturnPhrase []string `json:"returnPhrase"`
	ErrorCode string `json:"errorCode"`
	//Query        string   `json:"query"`
	Translation []string `json:"translation"`
	//Basic        struct {
	//	Phonetic   string
	//	UkPhonetic string
	//	UsPhonetic string
	//	UkSpeech   string
	//	UsSpeech   string
	//	explains   []string
	//} `json:"basic"`
	//Web []struct {
	//	Key   string
	//	Value []string
	//}
	//L    string `json:"l"`
	//Dict struct {
	//	Url string `json:"url"`
	//}
	//Webdict struct {
	//	Url string `json:"url"`
	//}
	//TSpeakUrl string
	//SpeakUrl  string
}
