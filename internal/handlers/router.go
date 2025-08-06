package handlers

import (
	"hospital-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	staffHandler := NewStaffHandler(db)
	patientHandler := NewPatientHandler(db)

	r.POST("/staff/create", staffHandler.CreateStaff)
	r.POST("/staff/login", staffHandler.Login)
	r.POST("/patient/search", middleware.AuthMiddleware(), patientHandler.SearchPatient)

	return r
}
