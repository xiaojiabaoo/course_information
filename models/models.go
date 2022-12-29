package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

var db *xorm.Engine

func InitMysql() error {
	var err error
	DsName := fmt.Sprintf("root:xiaoben@tcp(127.0.0.1:3306)/course_information?charset=utf8mb4")
	fmt.Println("数据库连接信息：" + DsName)
	db, err = xorm.NewEngine("mysql", DsName)
	if err != nil {
		return err
	}
	//todo 开发开启显示sql
	db.ShowSQL(false)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(240 * time.Second)
	//匹配表名
	db.Sync2()
	if err != nil {
		return err
	}
	return nil
}
