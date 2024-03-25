package main

import (
	"github.com/fast-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/fast_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("数据库错误")
	}

	db.AutoMigrate(
		&models.SysUser{},
		&models.SysRole{},
		&models.SysUserRole{},
		&models.SysMenu{},
		&models.SysRoleMenu{},
		&models.SysDictType{},
		&models.SysDictData{},
		&models.SysLoginLog{},
	)
}
