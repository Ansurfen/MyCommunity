package models

type Applications struct {
	First   string `json:"first" gorm:"varchar(110);not null"`
	Second  string `json:"second" gorm:"varchar(110);not null"`
	Third   string `json:"third" gorm:"varchar(110);not null"`
	Context string `json:"context" gorm:"varchar(1024);not null"`
	Status  int8   `json:"status" gorm:"not null"`
	Type    int8   `json:"type" gorm:"not null"`
	// Type + First + Second 校验是否重复发送
	Via string `json:"via" gorm:"varchar(255);unique;not null"`
}

func NewApplications() *Applications {
	return &Applications{}
}
