package common

import (
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) *models.User {
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return nil
	} else {
		user = tmpl.(models.User)
	}
	return &user
}
