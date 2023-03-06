package models

import "gorm.io/gorm"

type SiteSetting struct {
	gorm.Model
	NameEn string `gorm:"column:name_en;type:varchar(255);" json:"name_en"`
	NameZh string `gorm:"column:name;type:varchar(255);" json:"name_zh"`
	Value  string `gorm:"column:value;type:longtext;" json:"value"`
	Type   int    `gorm:"column:type;type:int(11);" json:"type"` // 1基础设置，2页脚徽标，3资料卡，4友链信息
}
