package question

import (
	dto "course_information/dto/question"
	"course_information/pkg/app"
	service "course_information/service/question"
	"github.com/gin-gonic/gin"
)

func GetQuestionData(c *gin.Context) {
	var (
		param      = dto.GetQuestionData{}
		err        error
		apiQuestion = service.ApiQuestion{}
	)
	err = app.BindAndValid(c, &param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	data, err := apiQuestion.GetQuestionData(param)
	if err != nil {
		app.Response(c, err, nil)
		return
	}
	app.Response(c, nil, data)
}
