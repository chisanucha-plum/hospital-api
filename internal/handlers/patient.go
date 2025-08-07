package handlers

import (
	"hospital-api/internal/models"
	"hospital-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
	return &PatientHandler{patientService: services.NewPatientService(db)}
}

func (h *PatientHandler) SearchPatient(c *gin.Context) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "Unauthorized: Hospital ID not found",
		})
		return
	}

	// get id parameter in URL path
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "ID parameter is required",
		})
		return
	}

	// Search by ID national_id or passport_id
	patient, err := h.patientService.SearchPatientByID(hospitalID.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to search patient: " + err.Error(),
		})
		return
	}

	if patient == nil {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Error:   "Patient not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Patient found",
		Data:    patient,
	})
}

func (h *PatientHandler) SearchPatients(c *gin.Context) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "Unauthorized: Hospital ID not found",
		})
		return
	}

	var req models.PatientSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	searchParams := map[string]string{
		"national_id":   req.NationalID,
		"patient_hn":    req.PatientHN,
		"first_name_th": req.FirstNameTH,
		"last_name_th":  req.LastNameTH,
		"first_name_en": req.FirstNameEN,
		"last_name_en":  req.LastNameEN,
		"passport_id":   req.PassportID,
		"phone_number":  req.PhoneNumber,
		"email":         req.Email,
	}

	patients, err := h.patientService.SearchPatients(hospitalID.(string), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to search patients: " + err.Error(),
		})
		return
	}

	if len(patients) == 0 {
		c.JSON(http.StatusNotFound, models.APIResponse{
			Success: false,
			Error:   "No patients found matching the criteria",
		})
		return
	}

	responseData := gin.H{
		"patients": patients,
		"count":    len(patients),
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Patients found",
		Data:    responseData,
	})
}
