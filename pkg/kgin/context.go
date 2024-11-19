package kgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"llm_service/internal/common"
	"net/http"
	"time"
)

type Context struct {
	*gin.Context
}

type Response struct {
	ErrorReason string            `json:"error_reason"`
	ErrorMsg    string            `json:"error_msg"`
	MetaData    map[string]string `json:"meta_data"`
	TraceId     string            `json:"trace_id"`
	ServerTime  int64             `json:"server_time"`
	Data        interface{}       `json:"data"`
	Code        int               `json:"code"`
}

func (c *Context) JsonOK(data any) {
	reply := Response{
		ErrorReason: "success",
		ErrorMsg:    "",
		MetaData:    nil,
		TraceId:     c.Request.Header.Get(common.TraceId),
		ServerTime:  time.Now().Unix(),
		Code:        http.StatusOK,
	}
	reply.Data = data
	c.JSON(200, reply)
}

func (c *Context) JsonError(code int, msg string) {
	reply := Response{
		ErrorReason: cast.ToString(code),
		ErrorMsg:    msg,
		TraceId:     c.Request.Header.Get(common.TraceId),
		ServerTime:  time.Now().Unix(),
		Data:        struct{}{},
		Code:        code,
	}
	c.JSON(200, reply)
}

func HandleFunc(handler func(c *Context)) func(*gin.Context) {
	h := func(c *gin.Context) {
		ctx := &Context{Context: c}
		traceId := c.GetHeader(common.TraceId)
		context.WithValue(c.Request.Context(), common.TraceId, traceId)
		handler(ctx)
	}
	return h
}
