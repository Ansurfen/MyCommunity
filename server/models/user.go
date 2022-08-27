package models

import (
	"MyCommunity/utils"
)

type User struct {
	Username    string `json:"username" gorm:"varchar(110);unique;not null"`
	Alias       string `json:"alias" gorm:"varchar(110);not null"`
	Password    string `json:"password" gorm:"size:255;not null"`
	Telephone   string `json:"telephone" gorm:"varchar(110);not null;"`
	Email       string `json:"email" gorm:"varchar(110);not null;"`
	School      string `json:"school" gorm:"varchar(120);"`
	Id          int64  `json:"id" gorm:"primary_key;not null"`
	Right       int8   `json:"right" gorm:"not null"`
	Communities string `json:"communities" gorm:"not null"`
	Profile     int8   `json:"profile" gorm:"not null"`
}

func NewUserWithPhone(username, telephone, password string) *User {
	return &User{
		Username:    username,
		Alias:       username,
		Password:    password,
		Telephone:   telephone,
		Email:       "",
		School:      "",
		Id:          utils.GenID(),
		Right:       None,
		Communities: "",
	}
}

func NewUserWithEmail(username, email, password string) *User {
	return &User{
		Username:    username,
		Alias:       username,
		Password:    password,
		Telephone:   "",
		Email:       email,
		School:      "",
		Id:          utils.GenID(),
		Right:       None,
		Communities: "",
	}
}
