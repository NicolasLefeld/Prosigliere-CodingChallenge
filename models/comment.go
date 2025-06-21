package models


type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}