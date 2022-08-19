package community

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	cname := ctx.PostForm("cname")
	db := sql.GetDB("general")
	var communities []models.Community
	db.Where("name LIKE ?", cname+"%").Find(&communities)
	res, err := json.Marshal(&communities)
	if err != nil {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
		return
	}
	common.SuccessRes(ctx, gin.H{"data": string(res)}, "获取数据成功")
}
