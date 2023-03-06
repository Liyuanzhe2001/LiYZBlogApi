package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	NickName    string `gorm:"column:nickname;type:varchar(255);" json:"nickname"`         // 昵称
	Description string `gorm:"column:description;type:varchar(255);" json:"description"`   // 描述
	Website     string `gorm:"column:website;type:varchar(255);" json:"website"`           // 站点
	Avatar      string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`             // 头像
	IsPublished string `gorm:"column:is_published;type:varchar(255);" json:"is_published"` // 公开或隐蔽
	Views       string `gorm:"column:views;type:int(11);" json:"views"`                    // 点击次数
}

func (table *Friend) TableName() string {
	return "friend"
}
