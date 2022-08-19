package models

type Posts struct {
	Id        string `json:"id" gorm:"primary_key;not null"`
	Title     string `json:"title" gorm:"not null"`
	Author    string `json:"author" gorm:"not null"`
	Timestamp string `json:"timestamp" gorm:"not null"`
	Context   string `json:"context" gorm:"not null"`
	Tags      string `json:"tags" gorm:"not null"`
	Comments  string `json:"comments" gorm:"type: longtext;not null"`
}

type Comment struct {
	Host      string `json:"host" gorm:"varchar(110);not null"`
	Context   string `json:"context" gorm:"not null"`
	Timestamp string `json:"timestamp" gorm:"not null"`
	Sub       string `json:"sub" gorm:"not null"`
}

type SubComment struct {
	From      string `json:"from" gorm:"not null"`
	To        string `json:"to" gorm:"not null"`
	Context   string `json:"context" gorm:"not null"`
	Timestamp string `json:"timestamp" gorm:"not null"`
}

func NewPosts(id, title, author, timestamp, context, tags string) *Posts {
	return &Posts{
		Id:        id,
		Title:     title,
		Author:    author,
		Timestamp: timestamp,
		Context:   context,
		Tags:      tags,
	}
}
