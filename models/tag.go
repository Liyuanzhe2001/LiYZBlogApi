package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string `gorm:"column:tag_name;type:varchar(255);" json:"tag_name"`
	Color   string `gorm:"column:color;type:varchar(255);" json:"color"` // 标签颜色(可选)
}
