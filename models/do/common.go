package do

type TSubjectChapter struct {
	TSubject  `xorm:"extends"`
	TChapter  `xorm:"extends"`
	UserOrder `xorm:"extends"`
}

type TChapterSection struct {
	TChapter `xorm:"extends"`
	TSection `xorm:"extends"`
	TPiece   `xorm:"extends"`
}

type TQuestionOption struct {
	TQuestion `xorm:"extends"`
	TOption   `xorm:"extends"`
}
