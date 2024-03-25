package models

type SysRole struct {
	BaseModel
	RoleName string `gorm:"not null;size:30;comment:角色名称"`
	RoleKey  string `gorm:"not null;size:100;comment:角色权限字符串"`
	RoleSort int    `gorm:"not null;size:4;type:int(4);comment:显示顺序"`
	Status   string `gorm:"not null;size:1;comment:角色状态（1正常 2停用）"`
	Remark   string `gorm:"size:500;default:'';comment:备注"`
}

func (SysRole) TableName() string {
	return "sys_role"
}
