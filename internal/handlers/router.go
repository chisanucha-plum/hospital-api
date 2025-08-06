package handlers

import (
	"hospital-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	staffHandler := NewStaffHandler(db)
	patientHandler := NewPatientHandler(db)

	r.POST("/staff/create", staffHandler.CreateStaff)
	r.POST("/staff/login", staffHandler.Login)
	r.GET("/patient/search/:id", middleware.AuthMiddleware(), patientHandler.SearchPatient)
	r.GET("/patient/search", middleware.AuthMiddleware(), patientHandler.SearchPatients)

	return r
}
