package commit

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pass(ctx *gin.Context) {
	var data models.Applications
	if tmpl, b := ctx.Get("application"); !b {
		common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "申请不存在")
		return
	} else {
		data = tmpl.(models.Applications)
	}
	if data.Type != models.PASS {
		common.FailRes(ctx, nil, "状态码校验失败")
		return
	}
	db := sql.GetDB("general")
	switch data.Status {
	case models.NEW: // 审批新建的社团
		var aps models.Applications
		via := utils.MD5(data.First + data.Second + strconv.Itoa(int(models.NEW)))
		db.Where("via = ?", via).First(&aps)
		if aps.Via == "" {
			common.FailRes(ctx, nil, "错误许可")
			return
		}
		db.Where("via = ?", via).Model(&aps).Update("status", models.ACTIVE)
		var community models.Community
		community.Id = utils.GenIDByString()
		community.Name = aps.First
		community.Hostname = aps.Second
		community.Context = aps.Context
		community.Timestamp = utils.NowTimestampByString()
		community.Status = models.ACTIVE
		db.Insert(&community)
		common.SuccessRes(ctx, nil, "审批成功")
	case models.ADD: // 加入社团
		var community models.Community
		if tmpl, b := ctx.Get("community"); !b {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "社团不存在")
			return
		} else {
			community = tmpl.(models.Community)
		}
		if len(community.Members) == 0 {
			community.Members = "[]"
		}
		var memebers models.Members
		if err := json.Unmarshal([]byte(community.Members), &memebers); err != nil {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
			return
		}
		memebers = append(memebers, data.First)
		// 需要一个更强大的tostring
		if tmpl, err := json.Marshal(&memebers); err != nil {
			common.CommonRes(ctx, http.StatusUnprocessableEntity, nil, "系统错误")
			return
		} else {
			community.Members = string(tmpl)
		}
		db.Where("id = ?", community.Id).Model(&models.Community{}).Update("members", community.Members)
		via := utils.MD5(data.First + data.Second + strconv.Itoa(int(models.ADD)))
		db.Where("via = ?", via).Model(&models.Applications{}).Update("status", models.ACTIVE)
		common.SuccessRes(ctx, nil, "审批成功")
	default:
	}
}
