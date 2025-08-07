package models

import "time"

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type UserPatient struct {
	NationalID   string    `json:"national_id" gorm:"primaryKey" `
	PatientHN    string    `json:"patient_hn" gorm:"unique" `
	FirstNameTH  string    `json:"first_name_th"`
	MiddleNameTH string    `json:"middle_name_th,omitempty"`
	LastNameTH   string    `json:"last_name_th"`
	FirstNameEN  string    `json:"first_name_en"`
	MiddleNameEN string    `json:"middle_name_en,omitempty"`
	LastNameEN   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PassportID   string    `json:"passport_id,omitempty" gorm:"unique"`
	PhoneNumber  string    `json:"phone_number,omitempty"`
	Email        string    `json:"email,omitempty"`
	Gender       Gender    `json:"gender" gorm:"type:varchar(1)"`
	HospitalID   string    `json:"hospital_id"`
	Hospital     Hospital  `json:"-" gorm:"foreignKey:HospitalID"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type UserStaff struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"-"`
	HospitalID string    `json:"hospital_id"`
	Hospital   Hospital  `json:"-" gorm:"foreignKey:HospitalID"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
