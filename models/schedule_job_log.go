package models

import "gorm.io/gorm"

type ScheduleJobLog struct {
	gorm.Model
	JobId      int64  `gorm:"column:job_id;type:bigint(20);" json:"job_id"`             // 任务id
	BeanName   string `gorm:"column:bean_name;type:varchar(255);" json:"bean_name"`     // spring bean名称
	MethodName string `gorm:"column:method_name;type:varchar(255);" json:"method_name"` // 方法名
	Params     string `gorm:"column:params;type:varchar(255);" json:"params"`           // 参数
	Status     int8   `gorm:"column:status;type:tinyint(4);" json:"status"`             // 任务状态
	Error      string `gorm:"column:error;type:text;" json:"error"`                     // 异常信息
	Times      int    `gorm:"column:times;type:int(11);" json:"times"`                  // 耗时（单位：毫秒）
}

func (table *ScheduleJobLog) TableName() string {
	return "ssite_settingchedule_job_log"

}
