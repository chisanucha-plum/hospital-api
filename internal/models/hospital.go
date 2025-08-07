package models

import "time"

type Hospital struct {
	ID        string        `json:"id" gorm:"primaryKey" `
	Name      string        `json:"name" gorm:"unique" `
	Staff     []UserStaff   `gorm:"foreignKey:HospitalID" json:"-"`
	Patients  []UserPatient `gorm:"foreignKey:HospitalID" json:"-"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
}
