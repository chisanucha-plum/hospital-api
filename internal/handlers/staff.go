package handlers

import (
	"hospital-api/internal/models"
	"hospital-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type StaffHandler struct {
	db *sqlx.DB
}

func NewStaffHandler(db *sqlx.DB) *StaffHandler {
	return &StaffHandler{db: db}
}

func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var input struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		HospitalID int    `json:"hospital_id"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	_, err := h.db.Exec("INSERT INTO staff (username, password, hospital_id) VALUES ($1, $2, $3)",
		input.Username, string(hashedPassword), input.HospitalID)
	if err != nil {
		log.Printf("DB error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff created successfully"})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var input struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		HospitalID int    `json:"hospital"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var staff models.UserStaff
	err := h.db.Get(&staff, "SELECT * FROM staff WHERE username = $1 AND hospital_id = $2",
		input.Username, input.HospitalID)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := services.GenerateJWT(staff.ID, staff.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
