package models

import "gorm.io/gorm"

type IdempotencyKey struct {
	gorm.Model
	Key     string `gorm:"primaryKey;type:varchar(100)"`
	BillId  uint   `gorm:"not null"`
	SplitId uint   `gorm:"not null"`
}
