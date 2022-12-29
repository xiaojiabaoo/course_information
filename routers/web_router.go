package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebRouter(engine *gin.Engine) {
	//engine.Use(common.TranslationMiddleware())
	engine.Static("/static/css", "./configs/static/css")
	engine.Static("/static/js", "./configs/static/js")
	engine.Static("/static/image", "./configs/static/image")
	engine.Static("/subject", "./configs/static/html/subject")
	engine.Static("/about", "./configs/static/html/about")
	engine.Static("/activity", "./configs/static/html/activity")
	engine.Static("/integral", "./configs/static/html/integral")
	engine.Static("/topic", "./configs/static/html/topic")
	engine.LoadHTMLFiles(
		"./configs/static/html/main.html", "./configs/static/html/login.html",
		"./configs/static/html/404.html", "./configs/static/html/500.html")
	engine.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	engine.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	})
	engine.GET("/500", func(c *gin.Context) {
		c.HTML(http.StatusOK, "500.html", gin.H{})
	})
	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
}
