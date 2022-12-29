package subject

import (
	dto "course_information/dto/subject"
	"course_information/models/model"
	"course_information/pkg/app"
	commonSer "course_information/service/common"
	service "course_information/service/subject"
	"github.com/gin-gonic/gin"
)

func GetSubjectList(c *gin.Context) {
	var (
		param      = dto.GetSubjectList{}
		err        error
		data       = make([]model.GetSubjectList, 0, 0)
		apiSubject = service.ApiSubject{}
		apiCommon = commonSer.ApiCommon{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	param.UserId = info
	data, err = apiSubject.GetSubjectList(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, data)
}

func TrySubject(c *gin.Context) {
	var (
		param      = dto.TrySubject{}
		err        error
		apiSubject = service.ApiSubject{}
		apiCommon = commonSer.ApiCommon{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	param.UserId = info
	err = apiSubject.TrySubject(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, nil)
}

func GetSubjectDetail(c *gin.Context) {
	var (
		param      = dto.GetSubjectDesc{}
		err        error
		apiSubject = service.ApiSubject{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	data, err := apiSubject.GetSubjectDetail(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, data)
}
