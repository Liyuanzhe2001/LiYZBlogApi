package models

import "gorm.io/gorm"

type About struct {
	gorm.Model
	NameEn string `gorm:"column:name_en;type:varchar(255);" json:"name_en"`
	NameZh string `gorm:"column:name;type:varchar(100);" json:"name_zh"`
	Value  string `gorm:"column:value;type:longtext;" json:"value"`
}

// TableName 表名
func (table *About) TableName() string {
	return "about"
}
