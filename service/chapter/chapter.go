package chapter

import (
	dto "course_information/dto/chapter"
	"course_information/models"
	"course_information/models/do"
	"course_information/models/model"
	"course_information/pkg"
	myErr "course_information/pkg/error"
)

type Chapter interface {
	GetChapterList(dto.GetChapterList) ([]model.GetChapterList, error)
}

type ApiChapter struct{}

func (a ApiChapter) GetChapterList(param dto.GetChapterList) ([]model.GetChapterList, error) {
	var (
		err        error
		data       = make([]model.GetChapterList, 0, 0)
		list       = make([]do.TChapterSection, 0, 0)
		chapter    = models.ApiChapter{}
		sectionMap = make(map[int][]model.Section)
		pieceMap   = make(map[int][]model.Piece)
		temp1Map    = make(map[int]struct{})
		temp2Map    = make(map[int]struct{})
	)

	if param.SubjectId == pkg.EMPTY_INT {
		return data, myErr.CustomError(myErr.SubjectEmptyError)
	}
	list, err = chapter.GetChapterList(param)
	if err != nil {
		return data, err
	}

	for _, v := range list {
		pieceMap[v.TPiece.SectionId] = append(pieceMap[v.TPiece.SectionId], model.Piece{
			Id:            v.TPiece.Id,
			Name:          v.TPiece.Name,
			QuestionCount: v.TPiece.QuestionCount,
			SectionId:     v.TPiece.SectionId,
		})
	}

	for _, v := range list {
		if _, ok := temp2Map[v.TSection.Id]; ok {
			continue
		}
		sectionMap[v.TSection.ChapterId] = append(sectionMap[v.TSection.ChapterId], model.Section{
			Id:            v.TSection.Id,
			Sequence:      v.TSection.Sequence,
			Name:          v.TSection.Name,
			QuestionCount: v.TSection.QuestionCount,
			ChapterId:     v.TSection.ChapterId,
			Piece:         pieceMap[v.TSection.Id],
		})
		temp2Map[v.TSection.Id] = struct{}{}
	}
	for _, v := range list {
		if _, ok := temp1Map[v.TChapter.Id]; ok {
			continue
		}
		data = append(data, model.GetChapterList{
			Id:            v.TChapter.Id,
			Sequence:      v.TChapter.Sequence,
			Name:          v.TChapter.Name,
			QuestionCount: v.TChapter.QuestionCount,
			SubjectId:     v.TChapter.SubjectId,
			Sort:          v.TChapter.Sort,
			Section:       sectionMap[v.TChapter.Id],
		})
		temp1Map[v.TChapter.Id] = struct{}{}
	}
	return data, nil
}
