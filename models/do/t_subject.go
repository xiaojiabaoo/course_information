package do

type TSubject struct {
	KnowledgeTreeId int    `xorm:"not null pk autoincr INT"`      // 知识树ID
	SubjectId       int    `xorm:"default 0 comment('课程ID') INT"` // 课程ID
	SubjectName     string `xorm:"comment('课程名') VARCHAR(255)"`   // 课程名
	PackageId       int    `xorm:"default 0 comment('专业ID') INT"` // 专业ID
	Sort            int    `xorm:"default 0 comment('排序') INT"`   // 排序
	Desc            string `xorm:"default 0 comment('课程简介说明') LONGTEXT"`
}
