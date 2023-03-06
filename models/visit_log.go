package models

import "gorm.io/gorm"

type VisitLog struct {
	gorm.Model
	Uuid      string `gorm:"column:uuid;type:varchar(36);" json:"uuid"`               // 游客标识码
	Uri       string `gorm:"column:uri;type:varchar(255);" json:"uri"`                // 请求接口
	Method    string `gorm:"column:method;type:varchar(255);" json:"method"`          // 请求方式
	Param     string `gorm:"column:param;type:varchar(2000);" json:"param"`           // 请求参数
	Behavior  string `gorm:"column:behavior;type:varchar(255);" json:"behavior"`      // 访问行为
	Content   string `gorm:"column:content;type:varchar(255);" json:"content"`        // 访问内容
	Remark    string `gorm:"column:remark;type:varchar(255);" json:"remark"`          // 备注
	Ip        string `gorm:"column:ip;type:varchar(255);" json:"ip"`                  // ip
	IpSource  string `gorm:"column:ip_source;type:varchar(255);" json:"ip_source"`    // ip来源
	Os        string `gorm:"column:os;type:varchar(255);" json:"os"`                  // 操作系统
	Browser   string `gorm:"column:browser;type:varchar(255);" json:"browser"`        // 浏览器
	Times     int    `gorm:"column:times;type:int(11);" json:"times"`                 // 请求耗时(毫秒)
	UserAgent string `gorm:"column:user_agent;type:varchar(2000);" json:"user_agent"` // user-agent用户代理
}

func (table *VisitLog) TableName() string {
	return "visit_log"
}
