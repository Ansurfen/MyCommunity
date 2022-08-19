package server

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
	"MyCommunity/utils"
)

func init() {
	common.SetResponseMode(common.STANDARD)
	db := sql.GetDB("general")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Posts{})
	db.AutoMigrate(&models.Community{})
	db.AutoMigrate(&models.Applications{})
	root := models.NewUserWithEmail("root", utils.EncodeAESWithKey(utils.BackedKey, "ansurfen@gmail.com"), utils.HashPassword("root"))
	root.Right = models.ROOT
	user1 := models.NewUserWithEmail("user1", utils.EncodeAESWithKey(utils.BackedKey, "user1@gmail.com"), utils.HashPassword("root"))
	user2 := models.NewUserWithEmail("user2", utils.EncodeAESWithKey(utils.BackedKey, "user2@gmail.com"), utils.HashPassword("root"))
	user3 := models.NewUserWithEmail("user3", utils.EncodeAESWithKey(utils.BackedKey, "user3@gmail.com"), utils.HashPassword("root"))
	db.InsertMany(root, user1, user2, user3)
}
