package models

import "gorm.io/gorm"

type LoginLog struct {
	gorm.Model
	Username    string `gorm:"column:username;type:varchar(255);" json:"username"`       // 用户名称
	Ip          string `gorm:"column:ip;type:varchar(255);" json:"ip"`                   // ip
	IpSource    string `gorm:"column:ip_source;type:varchar(255);" json:"ip_source"`     // ip来源
	Os          string `gorm:"column:os;type:varchar(255);" json:"os"`                   // 操作系统
	Browser     string `gorm:"column:browser;type:varchar(255);" json:"browser"`         // 浏览器
	Status      byte   `gorm:"column:status;type:bit(1);" json:"status"`                 // 登录状态
	Description string `gorm:"column:description;type:varchar(255);" json:"description"` // 操作描述
	UserAgent   string `gorm:"column:user_agent;type:varchar(2000);" json:"user_agent"`  // user-agent用户代理
}

func (table *LoginLog) TableName() string {
	return "login_log"
}
