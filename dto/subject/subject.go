package subject

type GetSubjectList struct {
	UserId int `json:"user_id"` //用户ID，非客户端参数
	Type   int `json:"type"`    //查询类型：1.主页菜单栏；2.所有课程；3.我的课程
}

type TrySubject struct {
	SubjectId int `json:"subject_id"` //课程ID
	UserId    int `json:"user_id"`    //用户ID，非客户端参数
}

type GetSubjectDesc struct {
	SubjectId int `json:"subject_id"` //课程ID
}
