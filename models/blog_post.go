package models

type BlogPost struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}
