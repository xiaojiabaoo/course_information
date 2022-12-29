package do

type TSection struct {
	Id            int    `xorm:"not null pk autoincr INT"`              // 节ID
	Sequence      string `xorm:"comment('节顺序') VARCHAR(255)"`               // 节顺序
	Name          string `xorm:"comment('节名称') VARCHAR(255)"`               // 节名称
	QuestionCount int    `xorm:"default 0 comment('节问题总数') INT"`  // 节问题总数
	ChapterId     int    `xorm:"default 0 comment('所属章节ID') INT"` // 所属章节ID
}
