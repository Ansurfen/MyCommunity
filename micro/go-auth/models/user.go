package models

type User struct {
	Username    string `json:"username"`
	Alias       string `json:"alias"`
	Password    string `json:"password"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	School      string `json:"school"`
	Id          int64  `json:"id"`
	Right       int8   `json:"right"`
	Communities string `json:"communities"`
	Image       string `json:"image"`
}
