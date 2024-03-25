package models

type SysRoleMenu struct {
	RoleID uint `gorm:"primarykey;comment:角色ID"`
	MenuID uint `gorm:"primarykey;comment:按钮ID"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
