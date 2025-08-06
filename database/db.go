package database

import (
	"fmt"
	"hospital-api/internal/configs"
	"hospital-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		configs.Envs.DBHost,
		configs.Envs.DBUser,
		configs.Envs.DBPassword,
		configs.Envs.DBName,
		configs.Envs.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

func InitSchema(db *gorm.DB) error {
	// GORM's AutoMigrate will create or update tables based on the struct definitions.
	err := db.AutoMigrate(&models.Hospital{}, &models.UserStaff{}, &models.UserPatient{})
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %v", err)
	}
	return nil
}
