package test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	// "os"
	// "encoding/json"
)

var (
	host      = "itrans.xfyun.cn"
	Algorithm = "hmac-sha256"
	HttpProto = "HTTP/1.1"
	uri       = "/v2/its"
	url       = "http://itrans.xfyun.cn/v2/its"
	//在控制台-机器翻译-服务接口认证信息获取
	appid  = "c4c9b4d0"
	Secret = "NWU4MTZlZDE3MmM2N2FhYmQ2MWYyNDhj"
	apiKey = "c9905a7610551c601da93a9f55d8d0c5"
	//text   = "今天天气怎么样" //请填写需要翻译的文本
	CN = "cn"
	EN = "en"
)

func Xunfei(txt string) string { //生成client 参数为默认
	client := &http.Client{}
	//生成要访问的url

	//提交请求
	var data1 []byte = []byte(txt)
	param := map[string]interface{}{
		"common": map[string]interface{}{
			"app_id": appid, //appid 必须带上，只需第一帧发送
		},
		"business": map[string]interface{}{ //business 参数，只需一帧发送
			"from": EN, //源语种
			"to":   CN, //目标语种
		},
		"data": map[string]interface{}{
			"text": base64.StdEncoding.EncodeToString(data1),
		},
	}
	tt, _ := json.Marshal(param)
	jsons := string(tt)
	jsoninfos := strings.NewReader(jsons)
	reqest, err := http.NewRequest("POST", url, jsoninfos)
	if err != nil {
		panic(err)
	}
	//增加header选项
	reqest.Header.Set("Content-Type", "application/json")
	reqest.Header.Set("Host", host)
	reqest.Header.Set("Accept", "application/json,version=1.0")
	currentTime := time.Now().UTC().Format(time.RFC1123)
	reqest.Header.Set("Date", currentTime)
	digest := "SHA-256=" + signBody(string(tt))
	reqest.Header.Set("Digest", digest)
	// 根据请求头部内容，生成签名
	sign := generateSignature(host, currentTime, "POST", uri, HttpProto, digest, Secret)
	// 组装Authorization头部
	authHeader := fmt.Sprintf(`api_key="%s", algorithm="%s", headers="host date request-line digest", signature="%s"`, apiKey, Algorithm, sign)
	reqest.Header.Set("Authorization", authHeader)

	//处理返回结果
	response, _ := client.Do(reqest)
	fmt.Println(response, err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	fmt.Println(body, err)
	var ttsReply XunfeiReply
	err = json.Unmarshal(body, &ttsReply)
	return ttsReply.Data.TransResult.Dst
	//fmt.Println(string(string(body)))
}

func generateSignature(host, date, httpMethod, requestUri, httpProto, digest string, secret string) string {
	// 不是request-line的话，则以header名称,后跟ASCII冒号:和ASCII空格，再附加header值
	var signatureStr string
	if len(host) != 0 {
		signatureStr = "host: " + host + "\n"
	}
	signatureStr += "date: " + date + "\n"
	// 如果是request-line的话，则以 http_method request_uri http_proto
	signatureStr += httpMethod + " " + requestUri + " " + httpProto + "\n"
	signatureStr += "digest: " + digest
	return hmacsign(signatureStr, secret)
}
func hmacsign(data, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}
func signBody(data string) string {
	// 进行sha256签名
	//fmt.Println(data)
	sha := sha256.New()
	sha.Write([]byte(data))
	encodeData := sha.Sum(nil)
	// 经过base64转换
	return base64.StdEncoding.EncodeToString(encodeData)
}
