package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	cname := ctx.PostForm("cname")
	db := sql.GetDB("general")
	var community models.Community
	db.Where("name = ?", cname).Find(&community)
	res, err := json.Marshal(&community)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": string(res)}, "获取数据成功")
}

func InfoApplication(ctx *gin.Context) {
	// cname := ctx.PostForm("cname")
	// var community models.Community
	// db := sql.GetDB("general")
	// db.Where("name = ?", cname).First(&community)
	// if community.Name != "" {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "Community name exist")
	// 	return
	// }
	// var applications []models.Application
	// db.Where("to = ?", cname).Find(&applications)
	// if data, err := json.Marshal(applications); err != nil {
	// 	common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "System error")
	// 	return
	// } else {
	// 	common.SuccessRes(ctx, gin.H{
	// 		"data": data,
	// 	}, "")
	// }
}
