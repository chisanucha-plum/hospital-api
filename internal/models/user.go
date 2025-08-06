package models

type Gender string

const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

func (g Gender) String() bool {
	return g == GenderMale || g == GenderFemale
}

type UserPatient struct {
	HospitalID   int    `gorm:"column:hospital_id" json:"hospital_id"`
	FirstNameTH  string `gorm:"column:first_name_th" json:"first_name_th"`
	MiddleNameTH string `gorm:"column:middle_name_th" json:"middle_name_th"`
	LastNameTH   string `gorm:"column:last_name_th" json:"last_name_th"`
	FirstNameEN  string `gorm:"column:first_name_en" json:"first_name_en"`
	MiddleNameEN string `gorm:"column:middle_name_en" json:"middle_name_en"`
	LastNameEN   string `gorm:"column:last_name_en" json:"last_name_en"`
	DateOfBirth  string `gorm:"column:date_of_birth" json:"date_of_birth"`
	PatientHN    string `gorm:"unique;column:patient_hn" json:"patient_hn"`
	NationalID   string `gorm:"primaryKey;column:national_id" json:"national_id"`
	PassportID   string `gorm:"unique;column:passport_id" json:"passport_id"`
	PhoneNumber  string `gorm:"column:phone_number" json:"phone_number"`
	Email        string `gorm:"column:email" json:"email"`
	Gender       Gender `gorm:"column:gender" json:"gender"`
}

func (UserPatient) TableName() string {
	return "patients"
}

type UserStaff struct {
	ID         int    `gorm:"primaryKey;column:id" json:"id"`
	Username   string `gorm:"unique;column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	HospitalID int    `gorm:"column:hospital_id" json:"hospital_id"`
}

func (UserStaff) TableName() string {
	return "staff"
}
