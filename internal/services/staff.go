package services

import (
	"errors"
	"fmt"
	"hospital-api/internal/models"
	"log"

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

func (s *StaffService) CreateStaffByRequest(req *models.CreateStaffRequest) (*models.UserStaff, error) {
	var hospital models.Hospital
	log.Printf("Looking for hospital: '%s' (length: %d)", req.HospitalName, len(req.HospitalName))
	if err := s.db.Where("name = ?", req.HospitalName).First(&hospital).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("hospital not found: %s", req.HospitalName)
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	staff := &models.UserStaff{
		Username:   req.Username,
		Password:   string(hashedPassword),
		HospitalID: hospital.ID,
	}

	if err := s.db.Create(staff).Error; err != nil {
		return nil, fmt.Errorf("failed to create staff: %v", err)
	}

	return staff, nil
}

func (s *StaffService) LoginByRequest(req *models.LoginRequest) (*models.UserStaff, error) {
	var staff models.UserStaff
	result := s.db.Where("username = ?", req.Username).First(&staff)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("database error: %v", result.Error)
	}

	err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &staff, nil
}
