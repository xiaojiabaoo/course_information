package models

import (
	dto "course_information/dto/chapter"
	"course_information/models/do"
	"github.com/pkg/errors"
)

type Chapter interface {
	GetChapterList(dto.GetChapterList) ([]do.TChapterSection, error)
}

type ApiChapter struct{}

func (a ApiChapter) GetChapterList(param dto.GetChapterList) ([]do.TChapterSection, error) {
	section := make([]do.TChapterSection, 0, 0)
	err := db.Cols("").Table("t_chapter").Alias("tc").
		Join("LEFT", []string{"t_section", "ts"}, "tc.id = ts.chapter_id").
		Join("LEFT", []string{"t_piece", "tp"}, "tp.section_id = ts.id").
		Where("tc.subject_id=?",param.SubjectId).
		OrderBy("tc.sort").OrderBy("ts.id").OrderBy("tc.id").Find(&section)
	if err != nil {
		return section, errors.Wrap(err,"根据科目ID查询章节信息失败")
	}
	return section, nil
}
