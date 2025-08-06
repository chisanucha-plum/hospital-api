package services

import (
	"errors"
	"fmt"
	"hospital-api/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StaffService struct {
	db *gorm.DB
}

func NewStaffService(db *gorm.DB) *StaffService {
	return &StaffService{db: db}
}

func (s *StaffService) CreateStaff(staff *models.UserStaff) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(staff.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	staff.Password = string(hashedPassword)

	if err := s.db.Create(staff).Error; err != nil {
		return fmt.Errorf("failed to create staff: %v", err)
	}
	return nil
}

func (s *StaffService) Login(username, password string) (*models.UserStaff, error) {
	var staff models.UserStaff
	result := s.db.Where("username = ?", username).First(&staff)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("database error: %v", result.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &staff, nil
}
