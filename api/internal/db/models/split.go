package models

import "gorm.io/gorm"

type Split struct {
	gorm.Model
	Bill      uint    `gorm:"not null"`
	UserId    uint    `gorm:"not null"`
	ContactId uint    `gorm:"type:varchar(100)"`
	Amount    float64 `gorm:"not null"`
	status    string  `gorm:"type:varchar(20);default:'due'"` // due, paid, failed, refunded
}
