package models

type Community struct {
	Id        string `json:"id" gorm:"primary_key;not null"`
	Name      string `json:"name" gorm:"varchar(110);not null"`
	Context   string `json:"context" gorm:"varchar(1024);not null"`
	Tags      string `json:"tags" gorm:"varchar(255);not null"`
	Image     string `json:"image" gorm:"varchar(50);not null"`
	Timestamp string `json:"timestamp" gorm:"not null"`
	Hostname  string `json:"hostname" gorm:"not null"`
	Admins    string `json:"admins" gorm:"not null"`
	Members   string `json:"memebers" gorm:"not null"`
	Notes     string `json:"notes" gorm:"not null"`
	Posts     string `json:"posts" gorm:"type: longtext;not null"`
	Status    int8   `json:"status" gorm:"not null"`
}

type Notification struct {
	Title   string `json:"title"`
	Context string `json:"context"`
	Timestamp string `json:"timestamp"`
}

type Post struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
	Tags      string `json:"tags"`
	Context   string `json:"context"`
	Score     int8   `json:"score"`
}

type Tags []string
type Admins []string
type Members []string

/*
	用户申请的列表存哪里，肯定存用户的表单...

	申请加入
	Type add
	Form(First) dycoldwind
	To(Second) [社团名字] // 通过社团名字去得到申请的列表
	Introduction //介绍自己

	创建社团
	Type new
	CName(First) 计算机社团
	Hostname(Second) dycoldwind
	Detail //社团简介

	获取社团信息和申请信息 // 也是不需要入库
	Type info
	GetType(First) add / new / ...
	Username(Second) //请求的用户，可能得鉴权

	删除社团
	Type del
	CName(First) 计算机社团 //
	Username(Second) dycoldwind // 谁说要删除的，要去鉴定权限

	删帖子 // 删除帖子直接删呗，但为了统一化拿到数据统一 json化
	Type delPost
	CName(First)
	Username(Second)

	发表帖子
	Type newPost
	CName(First) // 社团名
	Username // 为什么不直接jwt
	Context xxx 内容

	添加通知
	Type addNote


	删除通知
	Type delNote

	退出社团

	已申请的状态

	1. 管理员登录实现
*/

func NewCommunity() *Community {
	return &Community{}
}
