package models

type UserPatient struct {
	ID           int    `db:"id" json:"id"`
	HospitalID   int    `db:"hospital_id" json:"hospital_id"`
	FirstNameTH  string `db:"first_name_th" json:"first_name_th"`
	MiddleNameTH string `db:"middle_name_th" json:"middle_name_th"`
	LastNameTH   string `db:"last_name_th" json:"last_name_th"`
	FirstNameEN  string `db:"first_name_en" json:"first_name_en"`
	MiddleNameEN string `db:"middle_name_en" json:"middle_name_en"`
	LastNameEN   string `db:"last_name_en" json:"last_name_en"`
	DateOfBirth  string `db:"date_of_birth" json:"date_of_birth"`
	PatientHN    string `db:"patient_hn" json:"patient_hn"`
	NationalID   string `db:"national_id" json:"national_id"`
	PassportID   string `db:"passport_id" json:"passport_id"`
	PhoneNumber  string `db:"phone_number" json:"phone_number"`
	Email        string `db:"email" json:"email"`
	Gender       string `db:"gender" json:"gender"`
}

type UserStaff struct {
	ID         int    `db:"id" json:"id"`
	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"password"`
	HospitalID int    `db:"hospital_id" json:"hospital_id"`
}
