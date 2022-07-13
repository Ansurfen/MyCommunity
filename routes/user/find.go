package user

import (
	"MyCommunity/common"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	
	common.SuccessRes(ctx, nil, "密码找回成功")
}
