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
	var req models.CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	staff, err := h.staffService.CreateStaffByRequest(&req)
	if err != nil {
		log.Printf("Service error: %v", err)
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Failed to create staff: " + err.Error(),
		})
		return
	}

	response := models.CreateStaffResponse{
		Message: "Staff created successfully",
		StaffID: staff.ID,
	}
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Staff created successfully",
		Data:    response,
	})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Success: false,
			Error:   "Invalid input: " + err.Error(),
		})
		return
	}

	staff, err := h.staffService.LoginByRequest(&req)
	if err != nil {
		log.Printf("Login service error for user '%s': %v", req.Username, err)
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			Success: false,
			Error:   "Invalid credentials",
		})
		return
	}

	token, err := services.GenerateJWT(int(staff.ID), staff.HospitalID)
	if err != nil {
		log.Printf("JWT error: %v", err)
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Failed to generate token",
		})
		return
	}

	response := models.LoginResponse{Token: token}
	c.JSON(http.StatusOK, models.APIResponse{
		Success: true,
		Message: "Login successful",
		Data:    response,
	})
}
