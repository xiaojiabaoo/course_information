package do

import "time"

type UserOrder struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	OrderNo        string    `xorm:"comment('订单号') VARCHAR(255)"`
	Type           int       `xorm:"comment('类型：1.试用；2.购买；3.试用退货；4.购买退货；5.购买退款；6.购买退货退款') INT"`
	IsPayment      int       `xorm:"default 0 comment('是否已付款：0.无需付款；1.否；2.是；试用的课程无需付款') INT"`
	Amount         float64   `xorm:"comment('付款金额，试用的课程无金额') decimal"`
	BuyType        int       `xorm:"default 0 comment('购买类型：0.试用；1.包月；2.包季；3.包年；4.永久；试用产品默认为0') INT"`
	IsValid        int       `xorm:"comment('是否在有效期内：1.是；2.否') INT"`
	ExpirationTime int       `xorm:"comment('过期时间') INT"`
	UserId         int       `xorm:"comment('用户ID') INT"`
	SubjectId      int       `xorm:"comment('课程ID') INT"`
	AddTime        int       `xorm:"comment('更新时间，与date_time时间相同') INT"`
	UpTime         int       `xorm:"comment('更新时间，与date_time时间相同') INT"`
	DateTime       time.Time `xorm:"DATETIME"`
}
