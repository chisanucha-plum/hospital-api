package router

import (
	"hospital-api/internal/handlers"
	"hospital-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	staffHandler := handlers.NewStaffHandler(db)
	patientHandler := handlers.NewPatientHandler(db)

	staffRoutes := r.Group("/staff")
	{
		staffRoutes.POST("/create", staffHandler.CreateStaff)
		staffRoutes.POST("/login", staffHandler.Login)
	}

	patientRoutes := r.Group("/patient")
	patientRoutes.Use(middleware.AuthMiddleware())
	{
		patientRoutes.GET("/search/:id", patientHandler.SearchPatient)
		patientRoutes.GET("/search", patientHandler.SearchPatients)
	}

	return r
}
