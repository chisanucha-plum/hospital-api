package handlers

import (
	"hospital-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler(db *sqlx.DB) *PatientHandler {
	return &PatientHandler{patientService: services.NewPatientService(db)}
}

func (h *PatientHandler) SearchPatient(c *gin.Context) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var params struct {
		NationalID  string `json:"national_id"`
		PassportID  string `json:"passport_id"`
		FirstName   string `json:"first_name"`
		MiddleName  string `json:"middle_name"`
		LastName    string `json:"last_name"`
		DateOfBirth string `json:"date_of_birth"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
	}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	queryParams := map[string]string{
		"national_id":    params.NationalID,
		"passport_id":    params.PassportID,
		"first_name_th":  params.FirstName,
		"middle_name_th": params.MiddleName,
		"last_name_th":   params.LastName,
		"date_of_birth":  params.DateOfBirth,
		"phone_number":   params.PhoneNumber,
		"email":          params.Email,
	}

	patients, err := h.patientService.SearchPatients(hospitalID.(int), queryParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patients"})
		return
	}

	c.JSON(http.StatusOK, patients)
}
