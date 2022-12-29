package common

import (
	dto "course_information/dto/common"
	"course_information/pkg/app"
	service "course_information/service/common"
	"github.com/gin-gonic/gin"
)

func Verification(c *gin.Context) {
	var (
		param = dto.Verification{}
		err   error
		apiCommon = service.ApiCommon{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	err = apiCommon.Verification(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, nil)
}

func Login(c *gin.Context) {
	var (
		param = dto.Login{}
		err   error
		apiCommon = service.ApiCommon{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	login, err := apiCommon.Login(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, login)
}
