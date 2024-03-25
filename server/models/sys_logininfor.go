package models

import "time"

type SysLoginLog struct {
	ID            uint      `gorm:"primarykey"`
	LoginName     string    `gorm:"size:50;default:'';comment:登录账号"`
	IpAddr        string    `gorm:"size:128;default:'';comment:登录IP地址"`
	LoginLocation string    `gorm:"size:255;default:'';comment:登录地点"`
	Browser       string    `gorm:"size:50;default:'';comment:浏览器类型"`
	Os            string    `gorm:"size:50;default:'';comment:操作系统"`
	Status        string    `gorm:"size:1;default:1;comment:登录状态（1成功 2失败）"`
	Msg           string    `gorm:"size:255;default:'';comment:提示消息"`
	LoginTime     time.Time `gorm:"comment:登录时间"`
}

func (SysLoginLog) TableName() string {
	return "sys_login_log"
}
