package routes

import (
	"go-auth/common"
	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

func EncodeWithAES(ctx *gin.Context) {
	data, key := ctx.PostForm("data"), ctx.PostForm("key")
	data = utils.EncodeAESWithKey(data, key)
	common.SuccessRes(ctx, gin.H{"data": data}, "")
}

func DecodeWithAES(ctx *gin.Context) {
	data, key := ctx.PostForm("data"), ctx.PostForm("key")
	data = utils.DecodeAESWithKey(data, key)
	common.SuccessRes(ctx, gin.H{"data": data}, "")
}
