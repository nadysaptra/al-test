package database

import (
	"gorm.io/gorm"

	"azure-labs-test/app/models"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Event{}, &models.Ticket{}, &models.Payment{})
}
