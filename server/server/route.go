package server

import (
	"MyCommunity/middlewares"
	"MyCommunity/models"
	"MyCommunity/routes/admin"
	"MyCommunity/routes/commit"
	"MyCommunity/routes/community"
	"MyCommunity/routes/post"
	"MyCommunity/routes/user"
)

func (s *SEngine) UseRouter() *SEngine {

	userRouter := s.Group("/user")
	{
		userRouter.POST("/login", user.Login)
		userRouter.POST("/register", middlewares.CheckUniqueUsername(), middlewares.CheckUniqueKey(), user.Register)
		userRouter.POST("/edit/info", middlewares.AuthJWT(), user.EditInfo)
		userRouter.POST("/edit/pwd", user.EditPwd)
		userRouter.POST("/find", user.Find)
		userRouter.DELETE("/del", user.Cancel)
		userRouter.GET("/info", middlewares.AuthJWT(), user.Info)
		userRouter.POST("/image/update", middlewares.AuthJWT(), user.Update)
	}

	communityRouter := s.Group("/community")
	{
		communityRouter.POST("/new", middlewares.AuthJWT(), community.New)
		communityRouter.POST("/info", community.Info)
		communityRouter.POST("/add", middlewares.AuthJWT(), community.Add)
		communityRouter.POST("/edit", middlewares.AuthJWT(), middlewares.CheckCommunityRight(models.Admin), community.Edit)
		communityRouter.POST("/del", middlewares.AuthJWT(), middlewares.CheckCommunityRight(models.ROOT), community.Del)
		communityRouter.POST("/search", community.Search)
		communityRouter.POST("/cancel", middlewares.AuthJWT(), middlewares.CheckCommunityMember(), community.Cancel)
		communityRouter.POST("/post/add", middlewares.AuthJWT(), middlewares.CheckCommunityMember(), community.AddPost)
		communityRouter.POST("/post/info", community.InfoPost)
	}

	postRouter := s.Group("/post")
	{
		postRouter.POST("/info", post.Info)
		postRouter.POST("/add", middlewares.AuthJWT(), post.Add)
		postRouter.POST("/append", middlewares.AuthJWT(), post.Reply)
		postRouter.POST("/del", middlewares.AuthJWT(), post.DelComment)
		postRouter.POST("/del/sub", middlewares.AuthJWT(), post.DelSubComment)
		postRouter.POST("/edit/", middlewares.AuthJWT(), post.EditComment)
		postRouter.POST("/edit/sub", middlewares.AuthJWT(), post.EditSubComment)
	}

	commitRouter := s.Group("/commit", middlewares.AuthJWT())
	{
		commitRouter.POST("/info", commit.CommunityInfo)
		commitRouter.POST("/info/user", commit.UserInfo)
		commitRouter.POST("/pass", middlewares.CheckRight(models.Admin), commit.Pass)
		commitRouter.POST("/reject", middlewares.CheckRight(models.Admin), commit.Reject)
		commitRouter.POST("/change", middlewares.CheckRight(models.ROOT), commit.ChangeRight)
	}

	// 后台
	s.GET("/admin/info", middlewares.CheckRight(models.ROOT), admin.Info)
	s.POST("/admin/add", middlewares.CheckRight(models.ROOT), admin.Add)
	s.POST("/admin/del", middlewares.CheckRight(models.ROOT), admin.Del)
	s.POST("/admin/find", middlewares.CheckRight(models.ROOT), admin.Find)

	return s
}
