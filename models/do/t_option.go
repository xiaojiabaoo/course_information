package do

type TOption struct {
	Id         int    `xorm:"not null pk autoincr INT"`               // 选项ID
	QuestionId int    `xorm:"default 0 comment('所属问题ID') INT"`  // 所属问题ID
	Sequence   int    `xorm:"default 0 comment('顺序') INT"`      // 顺序
	Title      string `xorm:"comment('选项') VARCHAR(255)"`                 // 选项
	Content    string `xorm:"comment('选项内容') VARCHAR(255)"`               // 选项内容
	Correct    int    `xorm:"default 0 comment('是否是正确答案') INT"` // 是否是正确答案
}
