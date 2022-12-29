package configs

import (
	"course_information/models"
)

func InitConfig() error {
	var (
		err error
	)
	//初始化Mysql配置
	err = models.InitMysql()
	if err != nil {
		return err
	}
	//初始化Redis服务
	err = models.InitRedis()
	if err != nil {
		return err
	}
	return nil
}
