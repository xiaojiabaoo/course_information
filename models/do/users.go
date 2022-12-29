package do

import "time"

type Users struct {
	Id            int       `xorm:"not null pk autoincr INT"`                       // 用户ID
	NickName      string    `xorm:"comment('用户昵称/账号，用于展示使用') VARCHAR(255)"`         // 用户昵称/账号，用于展示使用
	Logo          string    `xorm:"comment('用户头像') VARCHAR(255)"`                   // 用户头像
	AreaCode      string    `xorm:"comment('用户手机区号') VARCHAR(255)"`                 // 用户手机区号
	Phone         int       `xorm:"default 0 comment('用户手机号码/账号，用户登录时使用') INT"`     // 用户手机号码/账号，用户登录时使用
	Email         string    `xorm:"comment('用户邮箱地址/账号，用户登录时使用') VARCHAR(255)"`      // 用户邮箱地址/账号，用户登录时使用
	Password      string    `xorm:"comment('用户账号密码') VARCHAR(255)"`                 // 用户账号密码
	Area          string    `xorm:"comment('用户所属的地区') VARCHAR(255)"`                // 用户所属的地区
	AddTime       int       `xorm:"default 0 comment('添加时间/用户注册时间') INT"`           // 添加时间/用户注册时间
	UpdateTime    int       `xorm:"default 0 comment('更新时间，与date_time时间相同') INT"`   // 更新时间，与date_time时间相同
	DateTime      time.Time `xorm:"DATETIME"`                                       // 更新时间，与update_time时间相同
	AccountStatus int       `xorm:"default 0 comment('账号状态：0.正常；1.黑名单；2.冻结；') INT"` // 账号状态：0.正常；1.黑名单；2.冻结；
}
