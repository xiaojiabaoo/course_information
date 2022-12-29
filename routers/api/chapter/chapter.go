package chapter

import (
	dto "course_information/dto/chapter"
	"course_information/pkg/app"
	service "course_information/service/chapter"
	"github.com/gin-gonic/gin"
)

func GetChapterList(c *gin.Context) {
	var (
		param      = dto.GetChapterList{}
		err        error
		apiChapter = service.ApiChapter{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	data, err := apiChapter.GetChapterList(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, data)
}
