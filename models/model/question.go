package model

type GetQuestionData struct {
	Id                int      `json:"id"`
	PieceId           int      `json:"piece_id"`
	Sequence          int      `json:"sequence"`
	QuestionType      string   `json:"question_type"`
	QuestionSource    string   `json:"question_source"`
	QuestionContent   string   `json:"question_content"`
	QuestionAnswer    string   `json:"question_answer"`
	Score             string   `json:"score"`
	MainNodeId        int      `json:"main_node_id"`
	MainNodeName      string   `json:"main_node_name"`
	MainNodeFrequency string   `json:"main_node_frequency"`
	Favorite          int      `json:"favorite"`
	CanPhoto          int      `json:"can_photo"`
	AvgCorrectRate    string   `json:"avg_correct_rate"`
	AnalysisType      int      `json:"analysis_type"`
	Analysis          string   `json:"analysis"`
	Option            []Option `json:"option"`
}

type Option struct {
	Id         int    `json:"id"`
	QuestionId int    `json:"question_id"`
	Sequence   int    `json:"sequence"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Correct    int    `json:"correct"`
}
