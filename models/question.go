package models

import (
	"course_information/models/do"
	"github.com/pkg/errors"
)

type Question interface {
	GetQuestionData(pieceId int) ([]do.TQuestionOption, error)
}

type ApiQuestion struct{}

func (a ApiQuestion) GetQuestionData(pieceId int) ([]do.TQuestionOption, error) {
	data := make([]do.TQuestionOption, 0, 0)
	err := db.Cols("tq.id", "tq.piece_id", "tq.sequence", "tq.question_type", "tq.question_source",
		"tq.question_content", "tq.question_answer", "tq.score", "tq.main_node_id", "tq.main_node_name",
		"tq.main_node_frequency", "tq.favorite", "tq.can_photo", "tq.avg_correct_rate", "tq.analysis_type",
		"tq.analysis", "to.id", "to.question_id", "to.sequence", "to.title", "to.content", "to.correct").
		Table("t_question").Alias("tq").
		Join("LEFT", []string{"t_option", "to"}, "tq.id = to.question_id").
		Where("tq.piece_id=?", pieceId).OrderBy("tq.sequence").OrderBy("to.sequence").Find(&data)
	if err != nil {
		return data, errors.Wrap(err, "根据块ID获取题目列表失败")
	}
	return data, nil
}
