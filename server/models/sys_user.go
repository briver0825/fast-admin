package models

import (
	"time"
)

type SysUser struct {
	BaseModel
	LoginName     string    `gorm:"not null;size:30;comment:登录账号"`
	UserName      string    `gorm:"size:30;default:'';comment:用户名"`
	UserType      string    `gorm:"size:2;default:00;comment:用户类型（00系统用户 01注册用户）"`
	Email         string    `gorm:"size:50;default:'';comment:用户邮箱"`
	Phone         string    `gorm:"size:11;default:'';comment:手机号"`
	Gender        string    `gorm:"size:1;default:1;comment:性别(0未知 1男 2女)"`
	Avatar        string    `gorm:"size:100;default:'';comment:头像路径"`
	Password      string    `gorm:"size:50;default:'';comment:密码"`
	Salt          string    `gorm:"size:20;default:'';comment:盐加密"`
	Status        string    `gorm:"size:1;default:1;comment:状态(1正常 2禁用)"`
	LoginIp       string    `gorm:"size:128;default:'';comment:最后登录ip"`
	LoginDate     time.Time `gorm:"comment:最后登录时间"`
	PwdUpdateDate time.Time `gorm:"comment:最后修改密码时间"`
	Remark        string    `gorm:"size:500;comment:备注;default:''"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
