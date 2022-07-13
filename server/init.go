package server

import (
	"MyCommunity/common"
	"MyCommunity/db/sql"
	"MyCommunity/models"
)

func init() {
	common.SetResponseMode(common.STANDARD)
	db := sql.GetDB("general")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Community{})
	db.AutoMigrate(&models.CommunityMember{})
	root := models.NewUserWithEmail("root", "ansurfen@gmail.com", "root")
	root.Right = models.ROOT
	user1 := models.NewUserWithEmail("user1", "user1@gmail.com", "root")
	user2 := models.NewUserWithEmail("user2", "user2@gmail.com", "root")
	user3 := models.NewUserWithEmail("user3", "user3@gmail.com", "root")
	db.InsertMany(root, user1, user2, user3)
}
