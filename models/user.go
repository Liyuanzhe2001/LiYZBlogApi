package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(255);" json:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(255);" json:"password"` // 密码
	Nickname string `gorm:"column:nickname;type:varchar(255);" json:"nickname"` // 昵称
	Avatar   string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`     // 头像地址
	Email    string `gorm:"column:email;type:varchar(255);" json:"email"`       // 邮箱
	Role     string `gorm:"column:role;type:varchar(255);" json:"role"`         // 角色访问权限
}

func (table *User) TableName() string {
	return "user"
}
