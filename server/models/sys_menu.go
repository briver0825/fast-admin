package models

type SysMenu struct {
	BaseModel
	MenuName  string `gorm:"not null;size:50;comment:菜单名称"`
	ParentID  uint   `gorm:"default:0;comment:父菜单ID"`
	OrderNum  int    `gorm:"type:int(4);default:0;comment:显示顺序"`
	Url       string `gorm:"size:200;default:#;comment:请求地址"`
	Target    string `gorm:"size:20;default:'';comment:打开方式（menuItem页签 menuBlank新窗口）"`
	MenuType  string `gorm:"size:1;default:'';comment:菜单类型（M目录 C菜单 F按钮）"`
	Visible   string `gorm:"size:1;default:1;comment:菜单状态（1显示 2隐藏）"`
	IsRefresh string `gorm:"size:1;default:1;comment:是否刷新（1刷新 2不刷新）"`
	Perms     string `gorm:"size:100;comment:权限标识"`
	Icon      string `gorm:"size:100;default:#;comment:菜单图标"`
	Remark    string `gorm:"size:500;default:'';comment:备注"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
