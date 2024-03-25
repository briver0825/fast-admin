package models

type SysDictData struct {
	BaseModel
	DictSort  int    `gorm:"type:int(4);default:0;comment:字典排序"`
	DictLabel string `gorm:"size:100;default:'';comment:字典标签"`
	DictValue string `gorm:"size:100;default:'';comment:字典键值"`
	DictType  string `gorm:"size:100;default:'';comment:字典类型"`
	CssClass  string `gorm:"size:100;default:'';comment:样式属性（其他样式扩展）"`
	ListClass string `gorm:"size:100;default:'';comment:表格回显样式"`
	IsDefault string `gorm:"size:1;default:'N';comment:是否默认（Y是 N否）"`
	Status    string `gorm:"size:1;default:1;comment:状态(1正常 2停用)"`
	Remarks   string `gorm:"size:500;default:'';comment:备注"`
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
