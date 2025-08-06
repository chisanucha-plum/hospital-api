package services

import (
	"encoding/json"
	"fmt"
	"hospital-api/internal/configs"
	"hospital-api/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type PatientService struct {
	db     *sqlx.DB
	client *http.Client
}

func NewPatientService(db *sqlx.DB) *PatientService {
	return &PatientService{
		db: db,
		client: &http.Client{
			Timeout: time.Duration(configs.Envs.HospitalAApiTimeout) * time.Second,
		},
	}
}

func (s *PatientService) SearchPatientByID(id string) (*models.UserPatient, error) {
	url := fmt.Sprintf("%s/patient/search/%s", configs.Envs.HospitalAApiUrl, id)
	resp, err := s.client.Get(url)
	if err != nil {
		log.Printf("Failed to call Hospital A API: %v", err)
		return nil, nil // ส่งคืน nil เพื่อให้ค้นจาก database
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Hospital A API returned status: %d", resp.StatusCode)
		return nil, nil
	}

	var patient models.UserPatient
	if err := json.NewDecoder(resp.Body).Decode(&patient); err != nil {
		log.Printf("Failed to decode Hospital A API response: %v", err)
		return nil, nil
	}

	patient.HospitalID = configs.Envs.HospitalID
	return &patient, nil
}

func (s *PatientService) SearchPatients(hospitalID int, params map[string]string) ([]models.UserPatient, error) {
	// Try searching Hospital A API by ID type
	if patient := s.searchPatientByIDTypes(hospitalID, params); patient != nil {
		return []models.UserPatient{*patient}, nil
	}

	// Build query and args for database search
	query, args := buildPatientQuery(hospitalID, params)

	var patients []models.UserPatient
	rows, err := s.db.NamedQuery(query, args)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var patient models.UserPatient
		if err := rows.StructScan(&patient); err != nil {
			return nil, fmt.Errorf("failed to scan patient: %v", err)
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

// searchPatientByIDTypes tries to find a patient by national_id or passport_id using Hospital A API
func (s *PatientService) searchPatientByIDTypes(hospitalID int, params map[string]string) *models.UserPatient {
	idKeys := []string{"national_id", "passport_id"}
	for _, key := range idKeys {
		if id, ok := params[key]; ok && id != "" {
			patient, err := s.SearchPatientByID(id)
			if err == nil && patient != nil && patient.HospitalID == hospitalID {
				return patient
			}
		}
	}
	return nil
}

// buildPatientQuery constructs the SQL query and arguments for patient search
func buildPatientQuery(hospitalID int, params map[string]string) (string, map[string]interface{}) {
	query := "SELECT * FROM patients WHERE hospital_id = :hospital_id"
	args := map[string]interface{}{"hospital_id": hospitalID}
	for key, value := range params {
		if value != "" {
			query += " AND " + key + " = :" + key
			args[key] = value
		}
	}
	return query, args
}
