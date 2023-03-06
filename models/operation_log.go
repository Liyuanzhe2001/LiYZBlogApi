package models

import "gorm.io/gorm"

type OperationLog struct {
	gorm.Model
	Username    string `gorm:"column:username;type:varchar(255);" json:"username"`       // 操作者用户名
	Uri         string `gorm:"column:uri;type:varchar(255);" json:"uri"`                 // 请求接口
	Method      string `gorm:"column:method;type:varchar(255);" json:"method"`           // 请求方式
	Param       string `gorm:"column:param;type:varchar(2000);" json:"param"`            // 请求参数
	Description string `gorm:"column:description;type:varchar(255);" json:"description"` // 操作描述
	Ip          string `gorm:"column:ip;type:varchar(255);" json:"ip"`                   // ip
	IpSource    string `gorm:"column:ip_source;type:varchar(255);" json:"ip_source"`     // ip来源
	Os          string `gorm:"column:os;type:varchar(255);" json:"os"`                   // 操作系统
	Browser     string `gorm:"column:browser;type:varchar(255);" json:"browser"`         // 浏览器
	Times       string `gorm:"column:times;type:int(11);" json:"times"`                  // 请求耗时（毫秒）
	UserAgent   string `gorm:"column:user_agent;type:varchar(2000);" json:"user_agent"`  // user-agent用户代理
}

func (table *OperationLog) TableName() string {
	return "operation_log"
}
