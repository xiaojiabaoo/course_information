package routers

import (
	"course_information/middleware"
	"course_information/routers/api/chapter"
	"course_information/routers/api/common"
	"course_information/routers/api/question"
	"course_information/routers/api/subject"
	"github.com/gin-gonic/gin"
)

func InitRouter() (engine *gin.Engine, err error) {
	engine = gin.Default()
	// 全局跨域支持
	engine.Use(middleware.CORSMiddleware())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	commons := engine.Group("/common")
	{
		commons.POST("/login", common.Login)         //用户登录接口（和注册接口共用）
		commons.POST("/verify", common.Verification) //验证码发送接口
	}
	subjects := engine.Group("/subject")
	{
		get := subjects.Group("/get")
		{
			get.POST("/list", subject.GetSubjectList)     //课程列表
			get.POST("/detail", subject.GetSubjectDetail) //课程详情
		}
		set := subjects.Group("/set")
		{
			set.POST("/try", subject.TrySubject) //课程试用
		}
	}
	chapters := engine.Group("/chapter")
	{
		get := chapters.Group("/get")
		{
			get.POST("/list", chapter.GetChapterList) //节块列表
		}
		/*set := sections.Group("/set")
		{

		}*/
	}
	questions := engine.Group("/question")
	{
		get := questions.Group("/get")
		{
			get.POST("/list", question.GetQuestionData)
		}
	}
	InitWebRouter(engine)
	return engine, err
}
