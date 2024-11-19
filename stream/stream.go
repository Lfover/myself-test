package main

import (
	"context"
	"encoding/json"
	"fmt"
	"myself-test/pkg/kgin"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// 设置HTTP头，告诉客户端这是一个流式响应
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	var c *kgin.Context
	ctx := context.WithValue(c.Context, "", time.Now())
	ch, err := p.au.SSECommon(ctx, appId, &req)
	if err != nil {
		return
	}
	// 使用Flusher接口来支持流式传输
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	for i := 0; i < 10; i++ {
		// 构造流数据
		fmt.Fprintf(w, "data: Message %d\n\n", i)
		// 刷新数据到客户端
		flusher.Flush()
		// 等待一段时间
		time.Sleep(1 * time.Second)
	}

	// 在结束时关闭连接
	fmt.Fprintf(w, "event: close\n")
	fmt.Fprintf(w, "data: closed\n\n")
	flusher.Flush()
}

func (c *Client) Send(ctx context.Context, server Server, prompt, message string, history []request.Message) (event <-chan dto.Data, err error) {
	if history == nil {
		c.log.WithContext(ctx).Infof("Azure4Gpt4: history 为空")
		history = make([]request.Message, 0)
	}
	if prompt != "" {
		newMessage := request.Message{
			Role:    request.SystemRole,
			Content: prompt,
		}
		// 将 newMessage 插入到 history 的第一个位置
		history = append([]request.Message{newMessage}, history...)
	} else {
		c.log.WithContext(ctx).Warnf("Azure4Gpt4: prompt 为空")
	}
	if message != "" {
		history = append(history, request.Message{
			Role:    request.UserRole,
			Content: message,
		})
	} else {
		c.log.WithContext(ctx).Warnf("Azure4Gpt4: message 为空")
	}
	c.option.Messages = history
	body, err := c.option.GetJsonBody()
	if err != nil {
		c.log.WithContext(ctx).Warnf("Azure4Gpt4: c.option.GetJsonBody() error:%v ", err)
		return nil, err
	}
	eventChan, err := c.client.Send(ctx, llm.PostMethod, server.GetHeader(c.apiKey), body)
	if err != nil {
		return nil, err
	}

	go func() {
		defer func() {
			close(c.eventChan)
			if err := recover(); err != nil {
				c.log.WithContext(ctx).Errorf("Azure4Gpt4: error, go routine panic错误：%v\n %s", err, debug.Stack())
				return
			}
		}()
		promptToken, _ := 0, 0 //tokenizer.CalToken(message)
		totalToken := 0
		for data := range eventChan {
			if data.Error != nil {
				// 存在异常，直接默认输出error兜底
				c.eventChan <- dto.Data{
					ErrorCode: common.OvertimeError,
					ErrorMsg:  common.OvertimeErrorMsg,
					IsEnd:     true,
				}
				return
			}

			var result Result
			err := json.Unmarshal([]byte(data.Data), &result)
			if err != nil {
				c.log.WithContext(ctx).Errorf("Azure4Gpt4: realData json marshal error：realData: %s, %s", data.Data, err.Error())
				continue
			}
			currentToken := 0 //result.GetTokens()
			totalToken += currentToken
			c.eventChan <- dto.Data{
				ErrorCode:  0,
				ErrorMsg:   "",
				Id:         result.Id,
				Object:     result.Object,
				Created:    result.Created,
				SentenceId: 0,
				IsEnd:      result.GetIsEnd(),
				Result:     result.GetContent(),
				Model:      result.Model,
				Usage: dto.Usage{
					PromptTokens:     promptToken,
					CompletionTokens: currentToken,
					TotalTokens:      totalToken,
				},
			}
			c.log.WithContext(ctx).Infof("Azure4Gpt4 real_data: %s", data.Data)

		}
	}()
	return c.eventChan, nil
}
func (s *SSEAzure4) SseCommonChat(ctx context.Context, biz request.LlmBiz, chat dto.SSECommonChat) (eventCh <-chan dto.Data, err error) {
	apiKeys := s.confD.Llm.Azure.Gpt4ApiKeys
	s.log.WithContext(ctx).Infof("used gpt4_azure biz:%v chat_param: %v", biz, chat)
	if len(apiKeys) <= 0 {
		return nil, common.NewErr(common.ECode110050, errors.New("请检查账号配置"))
	}
	accountKeys := strings.Split(apiKeys[0], ":")
	if len(accountKeys) < 2 {
		return nil, common.NewErr(common.ECode110050, errors.New("请检查账号配置"))
	}
	options := []azure4.ClientOption{azure4.WithModel()}
	if chat.Temperature != 0 {
		options = append(options, azure4.WithTemperature(chat.Temperature))
	}
	// 配置
	openAiClient := azure4.NewClient(s.confD.Llm.Azure.Gpt4Url, time.Second*300, apiKeys[0], s.log, options...)
	ch, err := openAiClient.Send(ctx, azure4.AzureServer, chat.Prompt, chat.Message, chat.History)
	if err != nil {
		return nil, err
	}

	return ch, nil
}
