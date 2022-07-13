package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(ctx *gin.Context) {
	var data models.Community
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	db := sql.GetDB("general")
	var community models.Community
	db.Where("name = ?", data.Name).First(&community)
	if community.Name != "" {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community name exist")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	admins, err := json.Marshal(models.NewCommunityAdmin(user.Username))
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
		return
	}
	data.Admin = string(admins)
	data.Status = models.WAITING
	db.Insert(&data)
	common.SuccessRes(ctx, nil, "提交成功，等待审核")
}
