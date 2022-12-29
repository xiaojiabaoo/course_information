package middleware

import (
	"github.com/gin-gonic/gin"
)

type ResponseCode int

type Response struct {
	Ret  int64       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code int64, msg string) {
	resp := &Response{Ret: code, Msg: msg, Data: make([]int, 0)}
	c.JSON(200, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	if data == "" || data == nil {
		data = make([]int, 0)
	}
	resp := &Response{Ret: 0, Msg: "操作成功", Data: data}
	c.JSON(200, resp)
	c.Abort()
}
