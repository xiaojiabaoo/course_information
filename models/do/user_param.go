package do

import "time"

type UserParam struct {
	Id       int       `xorm:"not null pk autoincr INT"`
	UserId   int       `xorm:"comment('用户ID') INT"`
	TryCount int       `xorm:"comment('用户的试用课程次数') INT"`
	TryTime  int       `xorm:"comment('用户每次试用的时间，单位：天') INT"`
	AddTime  int       `xorm:"comment('更新时间，与date_time时间相同') INT"`
	UpTime   int       `xorm:"comment('更新时间，与date_time时间相同') INT"`
	DateTime time.Time `xorm:"DATETIME"`
}
