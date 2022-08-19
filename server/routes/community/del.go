package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Del(ctx *gin.Context) {
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	var community models.Community
	if tmpl, b := ctx.Get("community"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community don't exist")
		return
	} else {
		community = tmpl.(models.Community)
	}
	if community.Hostname != user.Username {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限删除社团")
		return
	}
	var aps models.Applications
	aps.Type = models.DEL
	aps.Status = models.WAITING
	aps.First = community.Name
	aps.Second = user.Username
	aps.Via = utils.MD5(aps.First + aps.Second + strconv.Itoa(int(aps.Type)))
	db := sql.GetDB("general")
	db.Insert(&db)
	common.SuccessRes(ctx, nil, "删除信息已发送")
}
