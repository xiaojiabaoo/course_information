package model

type GetSubjectList struct {
	KnowledgeTreeId int       `json:"knowledge_tree_id"`
	SubjectId       int       `json:"subject_id"`
	SubjectName     string    `json:"subject_name"`
	PackageId       int       `json:"package_id"`
	Sort            int       `json:"sort"`
	Chapter         []Chapter `json:"chapter"`
	Status          int       `json:"status"` //拥有状态：1.试用；2.已购
	Desc            string    `json:"desc"`
	Progress        string    `json:"progress"` //学习进度，Status等于1和2时才有值
}

type Chapter struct {
	Id            int    `json:"id"`
	Sequence      string `json:"sequence"`
	Name          string `json:"name"`
	QuestionCount int    `json:"question_count"`
	SubjectId     int    `json:"subject_id"`
}

type GetSubjectDetail struct {
	KnowledgeTreeId int       `json:"knowledge_tree_id"`
	SubjectId       int       `json:"subject_id"`
	SubjectName     string    `json:"subject_name"`
	PackageId       int       `json:"package_id"`
	Sort            int       `json:"sort"`
	Desc            string    `json:"desc"`
}
