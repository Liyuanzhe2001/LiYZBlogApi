package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title            string `gorm:"column:title;type:varchar(255);" json:"title"`                         // 文章标题
	FirstPicture     string `gorm:"column:first_picture;type:varchar(255);" json:"first_picture"`         // 文章首图，用于随机文章展示
	Content          string `gorm:"column:content;type:longtext;" json:"content"`                         // 文章正文
	Description      string `gorm:"column:description;type:longtext;" json:"description"`                 // 描述
	IsPublished      uint8  `gorm:"column:is_published;type:tinyint(1);" json:"is_published"`             // 公开或私密
	IsRecommend      uint8  `gorm:"column:is_recommend;type:tinyint(1);" json:"is_recommend"`             // 推荐开关
	IsAppreciation   uint8  `gorm:"column:is_appreciation;type:tinyint(1);" json:"is_appreciation"`       // 赞赏开关
	IsCommentEnabled uint8  `gorm:"column:is_comment_enabled;type:tinyint(1);" json:"is_comment_enabled"` // 评论开关
	Views            int    `gorm:"column:views;type:int(11);" json:"views"`                              // 浏览次数
	Words            int    `gorm:"column:words;type:int(11);" json:"words"`                              // 文章字数
	ReadTime         int    `gorm:"column:read_time;type:int(11);" json:"read_time"`                      // 阅读时长(分钟)
	CategoryId       int64  `gorm:"column:category_id;type:bigint(20);" json:"category_id"`               // 文章分类
	IsTop            uint8  `gorm:"column:is_top;type:tinyint(1);" json:"is_top"`                         // 是否置顶
	Password         string `gorm:"column:password;type:varchar(255);" json:"password"`                   // 密码保护
	UserId           int64  `gorm:"column:user_id;type:bigint(20);" json:"user_id"`                       // 文章作者

	User     *User     `gorm:"foreignKey:id;references:user_id"`     //文章作者(因为是个人博客，也可以不加作者字段，暂且加上)
	Category *Category `gorm:"foreignKey:id;references:category_id"` //文章分类
	tags     *[]Tag    `gorm:"foreignKey:"`                          //文章标签

	//CategoryBasic *CategoryBasic `gorm:"foreignKey:id;references:category_id"`
}

// TableName 表名
func (table *Blog) TableName() string {
	return "blog"
}
