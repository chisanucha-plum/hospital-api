package models

import "github.com/golang-jwt/jwt/v5"

type CreateStaffRequest struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalName string `json:"hospital_name" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CreateStaffResponse struct {
	Message string `json:"message"`
	StaffID uint   `json:"staff_id"`
}

type PatientSearchRequest struct {
	NationalID  string `json:"national_id,omitempty"`
	PatientHN   string `json:"patient_hn,omitempty"`
	FirstNameTH string `json:"first_name_th,omitempty"`
	LastNameTH  string `json:"last_name_th,omitempty"`
	FirstNameEN string `json:"first_name_en,omitempty"`
	LastNameEN  string `json:"last_name_en,omitempty"`
	PassportID  string `json:"passport_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// JWT Claims Model - Type-safe JWT claims structure
type JWTClaims struct {
	StaffID    int `json:"staff_id"`
	HospitalID int `json:"hospital_id"`
	jwt.RegisteredClaims
}
