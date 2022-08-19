package user

import (
	"MyCommunity/common"

	"github.com/gin-gonic/gin"
)

const savePath string = "./images/i.png"

func Update(ctx *gin.Context) {
	if file, err := ctx.FormFile("file"); err != nil {
		common.FailRes(ctx, nil, "上传失败")
	} else {
		ctx.SaveUploadedFile(file, savePath)
	}
	common.SuccessRes(ctx, nil, "上传成功")
}
