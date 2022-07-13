package common

import (
	"MyCommunity/utils"
	"regexp"
)

const (
	_EMAIL_PATTERN = `\w+@\w+(\.\w+)+`
)

func CheckEmail(email string) bool {
	b, e := regexp.MatchString(_EMAIL_PATTERN, email)
	utils.Panic(e)
	return b
}
