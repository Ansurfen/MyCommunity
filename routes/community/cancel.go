package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cancel(ctx *gin.Context) {
	cname := ctx.PostForm("name")
	var memebers models.CommunityMember
	db := sql.GetDB("general")
	db.Where("name = ?", cname).First(&memebers)
	if memebers.Name == "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	}
	
}
