package main

import (
	"course_information/configs"
	public "course_information/pkg/log"
	rou "course_information/routers"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := configs.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 初始化路由
	routers, err := rou.InitRouter()
	if err != nil {
		public.ZapLog.Error("初始化路由失败", zap.Error(err))
		return
	}
	// 运行
	_ = routers.Run(":3000")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}