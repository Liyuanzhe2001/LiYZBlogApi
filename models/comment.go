package models

import (
	"LiYZBlog/models/vo"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	NickName        string `gorm:"column:nickname;type:varchar(255);" json:"nickname"`                 // 昵称
	Email           string `gorm:"column:email;type:varchar(255);" json:"email"`                       // 邮箱
	Content         string `gorm:"column:content;type:varchar(255);" json:"content"`                   // 评论内容
	Avatar          string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`                     // 头像(图片路径)
	Ip              string `gorm:"column:ip;type:varchar(255);" json:"ip"`                             // 评论者ip地址
	IsPublished     byte   `gorm:"column:is_published;type:varchar(255);" json:"is_published"`         // 公开或回收站
	IsAdminComment  byte   `gorm:"column:is_admin_comment;type:bit(1);" json:"is_admin_comment"`       // 博主回复
	Page            int    `gorm:"column:page;type:int(11);" json:"page"`                              // 0普通文章，1关于我页面，2友链页面
	IsNotice        byte   `gorm:"column:is_notice;type:bit(1);" json:"is_notice"`                     // 接收邮件提醒
	BlogId          int64  `gorm:"column:blog_id;type:bigint(20);" json:"blog_id"`                     // 所属的文章
	ParentCommentId int64  `gorm:"column:parent_comment_id;type:bigint(20);" json:"parent_comment_id"` // 父评论id，-1为根评论
	Website         string `gorm:"column:website;type:varchar(255);" json:"website"`                   // 个人网站
	Qq              string `gorm:"column:qq;type:varchar(255);" json:"qq"`                             // 如果评论昵称为QQ号，则将昵称和头像置为QQ昵称和QQ头像，并将此字段置为QQ号备份

	blog          *vo.BlogIdAndTitle // 所属的文章
	replyComments *[]Comment         //回复该评论的评论
}

func (table *Comment) TableName() string {
	return "comment"
}
