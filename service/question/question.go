package question

import (
	dto "course_information/dto/question"
	"course_information/models"
	"course_information/models/do"
	"course_information/models/model"
	"course_information/pkg"
	myErr "course_information/pkg/error"
)

type Question interface {
	GetQuestionData(dto.GetQuestionData) ([]model.GetQuestionData, error)
}

type ApiQuestion struct{}

func (a ApiQuestion) GetQuestionData(param dto.GetQuestionData) ([]model.GetQuestionData, error) {
	var (
		err          error
		apiQuestion  = models.ApiQuestion{}
		data         = make([]model.GetQuestionData, 0, 0)
		questionData = make([]do.TQuestionOption, 0, 0)
		optionMap    = make(map[int][]model.Option)
		tempMap      = make(map[int]struct{})
	)
	if param.PieceId == pkg.EMPTY_INT {
		return data, myErr.CustomError(myErr.PieceEmptyError)
	}
	questionData, err = apiQuestion.GetQuestionData(param.PieceId)
	if err != nil {
		return data, err
	}
	for _, v := range questionData {
		optionMap[v.TQuestion.Id] = append(optionMap[v.TQuestion.Id], model.Option{
			Id:         v.TOption.Id,
			QuestionId: v.TOption.QuestionId,
			Sequence:   v.TOption.Sequence,
			Title:      v.TOption.Title,
			Content:    v.TOption.Content,
			Correct:    v.TOption.Correct,
		})
	}
	for _, v := range questionData {
		if _, ok := tempMap[v.TQuestion.Id]; ok {
			continue
		}
		data = append(data, model.GetQuestionData{
			Id:                v.TQuestion.Id,
			PieceId:           v.TQuestion.PieceId,
			Sequence:          v.TQuestion.Sequence,
			QuestionType:      v.TQuestion.QuestionType,
			QuestionSource:    v.TQuestion.QuestionSource,
			QuestionContent:   v.TQuestion.QuestionContent,
			QuestionAnswer:    v.TQuestion.QuestionAnswer,
			Score:             v.TQuestion.Score,
			MainNodeId:        v.TQuestion.MainNodeId,
			MainNodeName:      v.TQuestion.MainNodeName,
			MainNodeFrequency: v.TQuestion.MainNodeFrequency,
			Favorite:          v.TQuestion.Favorite,
			CanPhoto:          v.TQuestion.CanPhoto,
			AvgCorrectRate:    v.TQuestion.AvgCorrectRate,
			AnalysisType:      v.TQuestion.AnalysisType,
			Analysis:          v.TQuestion.Analysis,
			Option:            optionMap[v.TQuestion.Id],
		})
		tempMap[v.TQuestion.Id] = struct{}{}
	}
	return data, nil
}
