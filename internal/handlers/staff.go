package handlers

import (
	"hospital-api/internal/models"
	"hospital-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StaffHandler struct {
	staffService *services.StaffService
}

func NewStaffHandler(db *gorm.DB) *StaffHandler {
	return &StaffHandler{staffService: services.NewStaffService(db)}
}

func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var staff models.UserStaff
	if err := c.BindJSON(&staff); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.staffService.CreateStaff(&staff); err != nil {
		log.Printf("Service error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff created successfully"})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	staff, err := h.staffService.Login(input.Username, input.Password)
	if err != nil {
		log.Printf("Login service error for user '%s': %v", input.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(staff.ID, staff.HospitalID)
	if err != nil {
		log.Printf("JWT error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
