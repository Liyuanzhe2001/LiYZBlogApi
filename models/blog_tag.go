package models

type BlogTag struct {
	BlogId int64 `gorm:"column:blog_id;type:bigint(20);" json:"blog_id"`
	TagId  int64 `gorm:"column:tag_id;type:bigint(20);" json:"tag_id"`
}

// TableName 表名
func (table *BlogTag) TableName() string {
	return "blog_tag"
}
