package do

import "time"

type UserLoginLogs struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	UserId      int       `xorm:"default 0 comment('用户ID') INT"`
	LoginTime   int       `xorm:"default 0 comment('用户登录时间') INT"`
	LoginIp     string    `xorm:"comment('用户登录IP') VARCHAR(255)"`
	LoginArea   string    `xorm:"comment('用户登录地区') VARCHAR(255)"`
	LoginMethod int       `xorm:"comment('户登录方式：1.手机号码登录；2.邮箱地址登录；.......') VARCHAR(255)"`
	DateTime    time.Time `xorm:"DATETIME"`
}
