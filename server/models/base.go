package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	ID        uint         `gorm:"primarykey"`
	CreatedAt time.Time    `gorm:"comment:创建时间"`
	UpdatedAt time.Time    `gorm:"comment:更新时间"`
	DeletedAt sql.NullTime `gorm:"index;comment:删除时间"`
	CreateBy  string       `gorm:"size:64;comment:创建者;default:''"`
	UpdateBy  string       `gorm:"size:64;comment:更新者;default:''"`
}
