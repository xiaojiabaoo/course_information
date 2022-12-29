package models

import (
	dto "course_information/dto/subject"
	"course_information/models/do"
	"github.com/pkg/errors"
)

type Subject interface {
	GetSubjectData(userId, types int) ([]do.TSubjectChapter, error)
	GetUserTrySubList(dto.TrySubject) ([]do.UserOrder, error)
	AddTrySubject(do.UserOrder) error
	GetSubjectDetail(subjectId int) (*do.TSubject, error)
}

type ApiSubject struct{}

func (a *ApiSubject) GetSubjectData(userId, types int) ([]do.TSubjectChapter, error) {
	subject := make([]do.TSubjectChapter, 0, 0)
	join := db.Cols("ts.knowledge_tree_id", "ts.subject_id", "ts.subject_name", "ts.package_id", "ts.sort",
		"tc.id", "tc.sequence", "tc.name", "tc.question_count", "tc.subject_id", "uo.type").
		Table("t_subject").Alias("ts").
		Join("LEFT", []string{"t_chapter", "tc"}, "ts.subject_id = tc.subject_id").
		Join("LEFT", []string{"user_order", "uo"}, "uo.subject_id = tc.subject_id")
	if types == 1 || types == 3 {
		join = join.Where("uo.user_id=? and is_valid=?", userId, 1)
	}
	err := join.OrderBy("ts.sort").OrderBy("tc.sort").Find(&subject)
	if err != nil {
		return subject, errors.Wrap(err, "查询课程表失败")
	}
	return subject, nil
}

func (a *ApiSubject) GetUserTrySubList(param dto.TrySubject) ([]do.UserOrder, error) {
	order := make([]do.UserOrder, 0, 0)
	err := db.Where("user_id=?", param.UserId).And("type=?", 1).Find(&order)
	if err != nil {
		return order, errors.Wrap(err, "查询用户订单失败")
	}
	return order, nil
}

func (a *ApiSubject) AddTrySubject(order do.UserOrder) error {
	_, err := db.Insert(order)
	if err != nil {
		return errors.Wrap(err, "添加课程试用失败")
	}
	return nil
}

func (a *ApiSubject) GetSubjectDetail(subjectId int) (*do.TSubject, error) {
	subject := &do.TSubject{}
	_, err := db.Where("subject_id=?", subjectId).Get(subject)
	if err != nil {
		return subject, errors.Wrap(err, "查询单个课程信息失败")
	}
	return subject, nil
}
