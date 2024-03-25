package models

type SysUserRole struct {
	UserID uint `gorm:"primarykey;comment:用户ID"`
	RoleID uint `gorm:"primarykey;comment:角色ID"`
}

func (SysUserRole) TableName() string {
	return "sys_user_role"
}
