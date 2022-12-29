package do

type TQuestion struct {
	Id                int    `xorm:"not null pk autoincr INT"` // 问题ID
	PieceId           int    `xorm:"default 0 comment('块ID') INT"` // 块ID
	Sequence          int    `xorm:"default 0 comment('顺序') INT"`  // 顺序
	QuestionType      string `xorm:"comment('问题类型') VARCHAR(255)"` // 问题类型
	QuestionSource    string `xorm:"comment('问题源') VARCHAR(255)"`  // 问题源
	QuestionContent   string `xorm:"comment('问题内容') VARCHAR(255)"` // 问题内容
	QuestionAnswer    string `xorm:"comment('问题答案') VARCHAR(255)"` // 问题答案
	Score             string `xorm:"comment('分数') VARCHAR(255)"`
	MainNodeId        int    `xorm:"default 0 comment('主节点ID') INT"`
	MainNodeName      string `xorm:"comment('主节点名称') VARCHAR(255)"`
	MainNodeFrequency string `xorm:"comment('主节点') VARCHAR(255)"`
	Favorite          int    `xorm:"default 0 comment('类型') INT"`
	CanPhoto          int    `xorm:"default 0 comment('类型') INT"`
	AvgCorrectRate    string `xorm:"comment('类型') VARCHAR(255)"`
	AnalysisType      int    `xorm:"default 0 comment('解析类型') INT"` // 解析类型
	Analysis          string `xorm:"comment('解析描述') VARCHAR(255)"`  // 解析描述
}
