package models

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Uuid      string `gorm:"column:uuid;type:varchar(36);" json:"uuid"`               // 游客标识码
	Ip        string `gorm:"column:ip;type:varchar(255);" json:"ip"`                  // ip
	IpSource  string `gorm:"column:ip_source;type:varchar(255);" json:"ip_source"`    // ip来源
	Os        string `gorm:"column:os;type:varchar(255);" json:"os"`                  // 操作系统
	Browser   string `gorm:"column:browser;type:varchar(255);" json:"browser"`        // 浏览器
	Pv        int    `gorm:"column:pv;type:int(11);" json:"pv"`                       // 访问页数统计
	UserAgent string `gorm:"column:user_agent;type:varchar(2000);" json:"user_agent"` // user-agent用户代理
}

func (table *Visitor) TableName() string {
	return "visitor"
}
