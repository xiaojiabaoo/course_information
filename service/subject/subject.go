package subject

import (
	dto "course_information/dto/subject"
	"course_information/models"
	"course_information/models/do"
	"course_information/models/model"
	"course_information/pkg"
	myErr "course_information/pkg/error"
	"time"
)

type Subject interface {
	GetSubjectList(dto.GetSubjectList) ([]model.GetSubjectList, error)
	TrySubject(dto.TrySubject) error
	GetSubjectDetail(dto.GetSubjectDesc) (model.GetSubjectDetail, error)
}

type ApiSubject struct{}

func (a *ApiSubject) GetSubjectList(param dto.GetSubjectList) ([]model.GetSubjectList, error) {
	var (
		resp       = make([]model.GetSubjectList, 0, 0)
		err        error
		apiSubject = models.ApiSubject{}
		chapterMap = make(map[int][]do.TChapter)
		tempMap    = make(map[int]struct{})
		data       = make([]do.TSubjectChapter, 0, 0)
	)
	//查询课程表
	switch param.Type {
	case 1: //主页
		data, err = apiSubject.GetSubjectData(param.UserId, param.Type)
		if err != nil {
			return resp, err
		}
	case 2: //所有课程
		data, err = apiSubject.GetSubjectData(param.UserId, param.Type)
		if err != nil {
			return resp, err
		}
	case 3: //我的课程
		data, err = apiSubject.GetSubjectData(param.UserId, param.Type)
		if err != nil {
			return resp, err
		}
	}

	for _, v := range data {
		chapterMap[v.TChapter.SubjectId] = append(chapterMap[v.TChapter.SubjectId], v.TChapter)
	}

	for _, v1 := range data {
		if _, ok := tempMap[v1.TSubject.SubjectId]; ok {
			continue
		}
		chapters := chapterMap[v1.TSubject.SubjectId]
		chapter := make([]model.Chapter, 0, 0)
		for _, v2 := range chapters {
			chapter = append(chapter, model.Chapter{
				Id:            v2.Id,
				Sequence:      v2.Sequence,
				Name:          v2.Name,
				QuestionCount: v2.QuestionCount,
				SubjectId:     v2.SubjectId,
			})
		}
		resp = append(resp, model.GetSubjectList{
			KnowledgeTreeId: v1.TSubject.KnowledgeTreeId,
			SubjectId:       v1.TSubject.SubjectId,
			SubjectName:     v1.TSubject.SubjectName,
			PackageId:       v1.TSubject.PackageId,
			Sort:            v1.TSubject.Sort,
			Chapter:         chapter,
			Status:          v1.UserOrder.Type,
		})
		tempMap[v1.TSubject.SubjectId] = struct{}{}
	}
	return resp, nil
}

func (a *ApiSubject) TrySubject(param dto.TrySubject) error {
	var (
		now        = time.Now()
		apiSubject = models.ApiSubject{}
		apiUser    = models.ApiUser{}
		err        error
		userParam  = &do.UserParam{}
		list       = make([]do.UserOrder, 0, 0)
		number     = pkg.CreateSerialNumber()
	)
	if param.SubjectId == 0 {
		return myErr.CustomError(myErr.SubjectEmptyError)
	}
	//获取到用户的可试用次数和使用时间
	userParam, err = apiUser.GetUserParam(param.UserId)
	if err != nil {
		return err
	}
	//获取用户的所有试用过的产品 包括正在试用的和过期的
	list, err = apiSubject.GetUserTrySubList(param)
	if err != nil {
		return err
	}
	//检测试用次数是否已达到了最高
	if userParam.TryCount <= len(list) {
		return myErr.CustomError(myErr.UserTryMaxError)
	}
	//检测购买的课程是否正在正常试用中
	for _, v := range list {
		if v.IsValid == 1 && v.SubjectId == param.SubjectId {
			return myErr.CustomError(myErr.UserTrySameError)
		}
	}
	//获取试用到期时间
	date := now.AddDate(0, 0, userParam.TryTime)
	order := do.UserOrder{
		OrderNo:        number,
		Type:           1,
		IsPayment:      0,
		BuyType:        0,
		IsValid:        1,
		ExpirationTime: int(date.Unix()),
		UserId:         param.UserId,
		SubjectId:      param.SubjectId,
		AddTime:        int(now.Unix()),
		UpTime:         int(now.Unix()),
		DateTime:       now,
	}
	err = apiSubject.AddTrySubject(order)
	if err != nil {
		return err
	}
	return nil
}

func (a *ApiSubject) GetSubjectDetail(param dto.GetSubjectDesc) (model.GetSubjectDetail, error) {
	var (
		err        error
		detail     = model.GetSubjectDetail{}
		apiSubject = models.ApiSubject{}
		data       = &do.TSubject{}
	)
	if param.SubjectId == pkg.EMPTY_INT {
		return detail, myErr.CustomError(myErr.SubjectEmptyError)
	}
	data, err = apiSubject.GetSubjectDetail(param.SubjectId)
	if err != nil {
		return model.GetSubjectDetail{}, err
	}
	detail.KnowledgeTreeId = data.KnowledgeTreeId
	detail.SubjectId = data.SubjectId
	detail.SubjectName = data.SubjectName
	detail.PackageId = data.PackageId
	detail.Sort = data.Sort
	detail.Desc = data.Desc
	return detail, nil
}
