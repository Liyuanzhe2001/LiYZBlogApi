package models

import "gorm.io/gorm"

type Moment struct {
	gorm.Model
	Content     string `gorm:"column:content;type:longtext;" json:"content"`         // 动态内容
	Likes       int    `gorm:"column:likes;type:int(11);" json:"likes"`              // 点赞数量
	IsPublished byte   `gorm:"column:is_published;type:bit(1);" json:"is_published"` // 是否公开
}

func (table *Moment) TableName() string {
	return "moment"
}
