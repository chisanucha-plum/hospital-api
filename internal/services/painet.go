package services

import (
	"errors"
	"fmt"
	"hospital-api/internal/models"

	"gorm.io/gorm"
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{db: db}
}

func (s *PatientService) SearchPatientByID(hospitalID string, id string) (*models.UserPatient, error) {

	var patient models.UserPatient
	result := s.db.Where("hospital_id = ?", hospitalID).
		Where("national_id = ? OR passport_id = ?", id, id).
		First(&patient)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query database: %v", result.Error)
	}

	return &patient, nil
}

func (s *PatientService) SearchPatients(hospitalID string, params map[string]string) ([]models.UserPatient, error) {

	query := s.db.Where("hospital_id = ?", hospitalID)
	for key, value := range params {
		if value != "" {
			query = query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}

	var patients []models.UserPatient
	if err := query.Find(&patients).Error; err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}

	return patients, nil
}
