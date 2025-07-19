package models

import (
	"gorm.io/gorm"
)

type AuditLog struct {
	gorm.Model
	EventType string `gorm:"type:varchar(50)"`
	BillId    uint   `gorm:"not null"`
	SplitId   uint   `gorm:"not null"`
	UserId    uint   `gorm:"not null"`
	Metadata  string `gorm:"type:text"`
}
