package main

import (
	"github.com/s9swata/billSplit/internal/db/models"
	"github.com/s9swata/billSplit/internal/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Bill{}, &models.Split{}, &models.IdempotencyKey{}, &models.AuditLog{})
}
