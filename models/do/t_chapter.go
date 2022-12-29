package do

type TChapter struct {
	Id            int    `xorm:"not null pk autoincr INT"`    // 章ID
	Sequence      string `xorm:"comment('章顺序') VARCHAR(255)"` // 章顺序
	Name          string `xorm:"comment('章名称') VARCHAR(255)"` // 章名称
	QuestionCount int    `xorm:"comment('章问题总数') INT"`        // 章问题总数
	SubjectId     int    `xorm:"comment('所属课程ID') INT"`       // 所属课程ID
	Sort          int    `xorm:"comment('排序') INT"`
}
