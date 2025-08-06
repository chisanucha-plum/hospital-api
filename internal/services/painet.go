package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"hospital-api/internal/configs"
	"hospital-api/internal/models"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type PatientService struct {
	db     *gorm.DB
	client *http.Client
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{
		db: db,
		client: &http.Client{
			Timeout: time.Duration(configs.Envs.HospitalAApiTimeout) * time.Second,
		},
	}
}

func (s *PatientService) SearchPatientByID(hospitalID int, id string) (*models.UserPatient, error) {
	// ลองค้นหาจาก Hospital A API ก่อน
	path := fmt.Sprintf("%s/patient/search/%s", configs.Envs.HospitalAApiUrl, id)
	resp, err := s.client.Get(path)
	if err != nil {
		log.Printf("Failed to call Hospital A API: %v", err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			var patient models.UserPatient
			if err := json.NewDecoder(resp.Body).Decode(&patient); err == nil {
				patient.HospitalID = hospitalID
				return &patient, nil
			}
		}
	}

	// Search in the local database using GORM
	var patient models.UserPatient
	result := s.db.Where("hospital_id = ?", hospitalID).
		Where("national_id = ? OR passport_id = ?", id, id).
		First(&patient)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Not found is not an error
		}
		return nil, fmt.Errorf("failed to query database: %v", result.Error)
	}

	return &patient, nil
}

func (s *PatientService) SearchPatients(hospitalID int, params map[string]string) ([]models.UserPatient, error) {
	// Try searching Hospital A API by ID type
	if patient := s.searchPatientByIDTypes(hospitalID, params); patient != nil {
		return []models.UserPatient{*patient}, nil
	}

	// Build query and args for database search using GORM
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

// searchPatientByIDTypes tries to find a patient by national_id or passport_id using Hospital A API
func (s *PatientService) searchPatientByIDTypes(hospitalID int, params map[string]string) *models.UserPatient {
	idKeys := []string{"national_id", "passport_id"}
	for _, key := range idKeys {
		if id, ok := params[key]; ok && id != "" {
			patient, err := s.SearchPatientByID(hospitalID, id)
			if err == nil && patient != nil && patient.HospitalID == hospitalID {
				return patient
			}
		}
	}
	return nil
}
