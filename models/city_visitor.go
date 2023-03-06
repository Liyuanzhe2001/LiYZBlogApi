package models

type CityVisitor struct {
	City string `gorm:"column:city;type:varchar(255);" json:"city"` // 城市名称
	Uv   int    `gorm:"column:uv;type:int(11);" json:"uv"`          // 独立访客数量
}

// TableName 表名
func (table *CityVisitor) TableName() string {
	return "city_visitor"
}
