package do

import "time"

type TUserlog struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	ProjectName string    `xorm:"comment('项目名称') VARCHAR(255)"`                       // 项目名称
	Platform    string    `xorm:"comment('平台') VARCHAR(255)"`                         // 平台
	UserId      int       `xorm:"default 0 comment('用户ID') INT"`            // 用户ID
	UserToken   string    `xorm:"comment('用户token') VARCHAR(255)"`                    // 用户token
	UserInfo    string    `xorm:"comment('用户信息：身份|userId|Token|个推CID') VARCHAR(255)"` // 用户信息：身份|userId|Token|个推CID
	MsgContent  string    `xorm:"comment('日志内容') VARCHAR(255)"`                       // 日志内容
	DetailInfo  string    `xorm:"comment('详情信息') VARCHAR(255)"`                       // 详情信息
	CreateTime  time.Time `xorm:"DATETIME"`                                           // 创建时间
}
