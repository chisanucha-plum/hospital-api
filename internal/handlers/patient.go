package handlers

import (
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// get id parameter in URL path
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	// Search by ID national_id or passport_id
	patient, err := h.patientService.SearchPatientByID(hospitalID.(int), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patient"})
		return
	}

	if patient == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) SearchPatients(c *gin.Context) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var searchParams map[string]string
	if err := c.BindJSON(&searchParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	patients, err := h.patientService.SearchPatients(hospitalID.(int), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"patients": patients, "count": len(patients)})
}
