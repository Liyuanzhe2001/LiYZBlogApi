package models

import "gorm.io/gorm"

type ScheduleJob struct {
	gorm.Model
	BeanName   string `gorm:"column:bean_name;type:varchar(255);" json:"bean_name"`     // spring bean名称
	MethodName string `gorm:"column:method_name;type:varchar(255);" json:"method_name"` // 方法名
	Params     string `gorm:"column:params;type:varchar(255);" json:"params"`           // 参数
	Cron       string `gorm:"column:cron;type:varchar(255);" json:"cron"`               // cron表达式
	Status     int8   `gorm:"column:status;type:tinyint(4);" json:"status"`             // 任务状态
	Remark     string `gorm:"column:remark;type:varchar(255);" json:"remark"`           // 备注
}

func (table *ScheduleJob) TableName() string {
	return "schedule_job"
}
