package models

type SysDictType struct {
	BaseModel
	DictName string `gorm:"size:100;default:'';comment:字典名称"`
	DictType string `gorm:"size:100;unique;default:'';comment:字典类型"`
	Status   string `gorm:"size:1;default:1;comment:状态(1正常 2停用)"`
	Remark   string `gorm:"size:500;default:'';comment:备注"`
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
