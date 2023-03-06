package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"column:category_name;type:varchar(255);" json:"category_name"`
}

// TableName 表名
func (table *Category) TableName() string {
	return "category"
}
