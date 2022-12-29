package do

type TPiece struct {
	Id            int    `xorm:"not null pk autoincr INT"`             // 块ID
	Name          string `xorm:"comment('块名称') VARCHAR(255)"`              // 块名称
	QuestionCount int    `xorm:"default 0 comment('块问题总数') INT"` // 块问题总数
	SectionId     int    `xorm:"default 0 comment('所属节ID') INT"` // 所属节ID
}
