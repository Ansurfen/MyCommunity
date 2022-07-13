package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	name := ctx.PostForm("cname")
	var community models.Community
	db := sql.GetDB("general")
	db.Where("name = ?", name).First(&community)
	if community.Name == "" || community.Status == models.WAITING {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	// 还有重复提交的情况
	application := models.NewApplication(user.Username, name)
	db.Insert(application)
	common.SuccessRes(ctx, nil, "申请已经提交")
}
