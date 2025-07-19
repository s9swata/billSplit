package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100)"`
	Email       *string    `gorm:"type:varchar(100);unique"`
	Age         uint8      `gorm:"type:int"`
	Birthday    *time.Time `gorm:"type:date"`
	PhoneNumber *string    `gorm:"type:varchar(15);unique"`
	UpiId       *string    `gorm:"type:varchar(50);unique"`
	DeviceId    *string    `gorm:"type:varchar(100);unique"`
}
