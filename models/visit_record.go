package models

import "gorm.io/gorm"

type VisitRecord struct {
	gorm.Model
	Pv   int    `gorm:"column:pv;type:int(11);" json:"pv"`          // 访问量
	Uv   int    `gorm:"column:uv;type:int(11);" json:"uv"`          // 独立用户
	Date string `gorm:"column:date;type:varchar(255);" json:"date"` // 日期"02-23"
}

func (table *VisitRecord) TableName() string {
	return "visit_record"
}
