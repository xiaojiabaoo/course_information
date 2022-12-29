package app

import (
	errors "course_information/pkg/error"
	"course_information/pkg/log/zap"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var Log *zap.ZapLog

// Response setting gin.JSON
func Response(c *gin.Context, err error, data interface{}) {
	var (
		msg  string
		code int
	)
	if err != nil {
		if tem, ok := err.(*errors.MyError); ok {
			msg = errors.GetMsg(tem.Code())
			code = tem.Code()
		} else {
			msg = errors.GetMsg(errors.ServerError)
			code = errors.ServerError
			fmt.Println(err)
		}
	} else {
		msg = errors.GetMsg(errors.Success)
		code = errors.Success
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
	c.JSON(200, &Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	c.Abort()
}
