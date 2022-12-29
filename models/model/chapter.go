package model

type GetChapterList struct {
	Id            int       `json:"id"`
	Sequence      string    `json:"sequence"`
	Name          string    `json:"name"`
	QuestionCount int       `json:"question_count"`
	SubjectId     int       `json:"subject_id"`
	Sort          int       `json:"sort"`
	Section       []Section `json:"section"`
}

type Section struct {
	Id            int     `json:"id"`
	Sequence      string  `json:"sequence"`
	Name          string  `json:"name"`
	QuestionCount int     `json:"question_count"`
	ChapterId     int     `json:"chapter_id"`
	Piece         []Piece `json:"piece"`
}

type Piece struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	QuestionCount int    `json:"question_count"`
	SectionId     int    `json:"section_id"`
}
