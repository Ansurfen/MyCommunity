package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommunityInfo(ctx *gin.Context) {
	var data models.Applications
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Type != models.INFO {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	var user *models.User
	if user = common.GetUser(ctx); user == nil {
		return
	}
	var aps []models.Applications
	db := sql.GetDB("general")
	// get type
	if data.Status == models.ERR {
		common.FailRes(ctx, nil, "类型校验错误")
		return
	}
	//应该还有验权的
	sql := ""
	switch data.Status {
	case models.NEW:
		sql = "type = ? and second = ?"
		db.Where(sql, data.Status, user.Username).Find(&aps)
	case models.ALL:
		//得列举出所有情况呃
		if user.Right == models.ROOT || user.Right == models.Admin {
			sql = "type > 6 and type < 11"
			db.Where(sql).Find(&aps)
		} else {
			sql = "type > 6 and type < 11 and second = ?"
			db.Where(sql, user.Username).Find(&aps)
		}
	}
	res, err := json.Marshal(&aps)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": string(res)}, "获取信息成功")
}

func UserInfo(ctx *gin.Context) {
	var data models.Applications
	if err := ctx.ShouldBindJSON(&data); err != nil {
		common.CommonRes(ctx, http.StatusBadRequest, gin.H{"err": err.Error()}, "Fail to parser JSON")
		return
	}
	if data.Type != models.INFO {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	var user models.User
	if tmpl, b := ctx.Get("user"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "User don't exist")
		return
	} else {
		user = tmpl.(models.User)
	}
	if user.Right != models.ROOT {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "没有权限")
		return
	}
	var users []models.User
	db := sql.GetDB("general")
	db.Find(&users)
	for i := 0; i < len(users); i++ {
		if users[i].Telephone != "" {
			users[i].Telephone = utils.DecodeAESWithKey(utils.BackedKey, users[i].Telephone)
		}
		if users[i].Email != "" {
			users[i].Email = utils.DecodeAESWithKey(utils.BackedKey, users[i].Email)
		}
		users[i].Password = ""
	}
	res, err := json.Marshal(&users)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": string(res)}, "获取信息成功")
}
