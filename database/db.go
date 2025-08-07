package database

import (
	"fmt"
	"hospital-api/internal/configs"
	"hospital-api/internal/models"
	"log"

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
	log.Println("Initializing database schema...")

	// GORM's AutoMigrate
	err := db.AutoMigrate(&models.Hospital{}, &models.UserStaff{}, &models.UserPatient{})
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %v", err)
	}

	log.Println("Database schema initialized successfully")

	if err := SeedMockData(db); err != nil {
		log.Printf("Warning: Failed to seed mock data: %v", err)
	}

	GetMockDataSummary(db)

	return nil
}
