package models

type CensoredWords struct {
	ID   uint64 `json:"id" ;gorm:"primary_key"`
	Word string `json:"word" ;gorm:"column:word"`
}

type CreateCommentRequestBody struct {
	NewsId   uint64 `json:"news_id"`
	ParentId uint64 `json:"parent_id"`
	Text     string `json:"text"`
	UserId   uint64 `json:"user_id"`
	Censored bool
}
