package models

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	CreatorId          uint    `gorm:"not null"`
	TotalAmount        float64 `gorm:"not null"`
	MerchantVpa        *string `gorm:"type:varchar(50)"`
	Status             string  `gorm:"type:varchar(20);default:'open'"` //open, pending, settled, failed
	SettlementDeadline *time.Time
}
