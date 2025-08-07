package database

import (
	"hospital-api/internal/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedMockData mock  3 tables
func SeedMockData(db *gorm.DB) error {
	log.Println("Starting mock data seeding...")

	if hasValidMockData(db) {
		log.Println("Valid mock data already exists, skipping seed...")
		return nil
	}

	if err := clearExistingData(db); err != nil {
		return err
	}

	if err := seedHospitals(db); err != nil {
		return err
	}

	if err := seedStaff(db); err != nil {
		return err
	}

	if err := seedPatients(db); err != nil {
		return err
	}

	log.Println("Mock data seeding completed successfully!")
	return nil
}

func hasValidMockData(db *gorm.DB) bool {
	var hospitalCount, staffCount, patientCount int64

	db.Model(&models.Hospital{}).Count(&hospitalCount)
	db.Model(&models.UserStaff{}).Count(&staffCount)
	db.Model(&models.UserPatient{}).Count(&patientCount)

	return hospitalCount == 3 && staffCount == 3 && patientCount == 6
}

func clearExistingData(db *gorm.DB) error {
	log.Println("Clearing existing data...")

	if err := db.Where("1 = 1").Delete(&models.UserPatient{}).Error; err != nil {
		return err
	}
	if err := db.Where("1 = 1").Delete(&models.UserStaff{}).Error; err != nil {
		return err
	}
	if err := db.Where("1 = 1").Delete(&models.Hospital{}).Error; err != nil {
		return err
	}

	return nil
}

func seedHospitals(db *gorm.DB) error {
	log.Println("Seeding hospitals...")

	hospitals := []models.Hospital{
		{ID: 1, Name: "โรงพยาบาลศิริราช"},
		{ID: 2, Name: "โรงพยาบาลจุฬาลงกรณ์"},
		{ID: 3, Name: "โรงพยาบาลรามาธิบดี"},
	}

	for _, hospital := range hospitals {
		if err := db.Create(&hospital).Error; err != nil {
			log.Printf(" Error creating hospital %s: %v", hospital.Name, err)
			return err
		}
		log.Printf("Created hospital: %s (ID: %d)", hospital.Name, hospital.ID)
	}

	return nil
}

func seedStaff(db *gorm.DB) error {
	log.Println("Seeding staff...")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	staff := []models.UserStaff{
		{ID: 1, Username: "admin1", Password: string(hashedPassword), HospitalID: 1},
		{ID: 2, Username: "admin2", Password: string(hashedPassword), HospitalID: 2},
		{ID: 3, Username: "admin3", Password: string(hashedPassword), HospitalID: 3},
	}

	for _, s := range staff {
		if err := db.Create(&s).Error; err != nil {
			log.Printf("Error creating staff %s: %v", s.Username, err)
			return err
		}
		log.Printf("Created staff: %s (Hospital ID: %d)", s.Username, s.HospitalID)
	}

	return nil
}

func seedPatients(db *gorm.DB) error {
	log.Println("Seeding patients...")

	const dateFormat = "2006-01-02"

	// Mock birth dates
	birthDate1990, _ := time.Parse(dateFormat, "1990-01-15")
	birthDate1985, _ := time.Parse(dateFormat, "1985-06-20")
	birthDate1992, _ := time.Parse(dateFormat, "1992-12-10")

	patients := []models.UserPatient{
		{
			NationalID: "1234567890123", PatientHN: "HN001",
			FirstNameTH: "สมชาย", LastNameTH: "ใจดี",
			FirstNameEN: "Somchai", LastNameEN: "Jaidee",
			DateOfBirth: birthDate1990, PassportID: "P001",
			PhoneNumber: "081-234-5678", Email: "somchai@test.com",
			Gender: "M", HospitalID: "1",
		},
		{
			NationalID: "2345678901234", PatientHN: "HN002",
			FirstNameTH: "สมหญิง", LastNameTH: "สวยงาม",
			FirstNameEN: "Somying", LastNameEN: "Suaynam",
			DateOfBirth: birthDate1985, PassportID: "P002",
			PhoneNumber: "082-345-6789", Email: "somying@test.com",
			Gender: "F", HospitalID: "2",
		},
		{
			NationalID: "3456789012345", PatientHN: "HN003",
			FirstNameTH: "วิชัย", LastNameTH: "เก่งมาก",
			FirstNameEN: "Wichai", LastNameEN: "Kengmak",
			DateOfBirth: birthDate1992, PassportID: "P003",
			PhoneNumber: "083-456-7890", Email: "wichai@test.com",
			Gender: "M", HospitalID: "3",
		},
		{
			NationalID: "4567890123456", PatientHN: "HN004",
			FirstNameTH: "วันทนา", LastNameTH: "รักดี",
			FirstNameEN: "Wantana", LastNameEN: "Rakdee",
			DateOfBirth: birthDate1985, PassportID: "P004",
			PhoneNumber: "084-567-8901", Email: "wantana@test.com",
			Gender: "F", HospitalID: "1",
		},
		{
			NationalID: "5678901234567", PatientHN: "HN005",
			FirstNameTH: "ประพจน์", LastNameTH: "สมประสงค์",
			FirstNameEN: "Prapot", LastNameEN: "Somprasong",
			DateOfBirth: birthDate1992, PassportID: "P005",
			PhoneNumber: "085-678-9012", Email: "prapot@test.com",
			Gender: "M", HospitalID: "2",
		},
		{
			NationalID: "6789012345678", PatientHN: "HN006",
			FirstNameTH: "สุภาพร", LastNameTH: "เรียนดี",
			FirstNameEN: "Supaporn", LastNameEN: "Riandee",
			DateOfBirth: birthDate1990, PassportID: "P006",
			PhoneNumber: "086-789-0123", Email: "supaporn@test.com",
			Gender: "F", HospitalID: "3",
		},
	}

	for _, patient := range patients {
		if err := db.Create(&patient).Error; err != nil {
			log.Printf(" Error creating patient %s: %v", patient.FirstNameTH, err)
			return err
		}
		log.Printf(" Created patient: %s %s (HN: %s, Hospital ID: %s)",
			patient.FirstNameTH, patient.LastNameTH, patient.PatientHN, patient.HospitalID)
	}

	return nil
}

func GetMockDataSummary(db *gorm.DB) {
	var hospitalCount, staffCount, patientCount int64

	db.Model(&models.Hospital{}).Count(&hospitalCount)
	db.Model(&models.UserStaff{}).Count(&staffCount)
	db.Model(&models.UserPatient{}).Count(&patientCount)

	log.Println(" Mock Data Summary:")
	log.Printf(" Hospitals: %d", hospitalCount)
	log.Printf(" Staff: %d", staffCount)
	log.Printf(" Patients: %d", patientCount)

	if hospitalCount > 0 {
		log.Println("\n Hospitals in system:")
		var hospitals []models.Hospital
		db.Find(&hospitals)
		for _, h := range hospitals {
			log.Printf("  - ID: %d, Name: %s", h.ID, h.Name)
		}
	}

	if staffCount > 0 {
		log.Println("\n Staff credentials (password: password123):")
		var staff []models.UserStaff
		db.Preload("Hospital").Find(&staff)
		for _, s := range staff {
			log.Printf("  - Username: %s, Hospital: %s", s.Username, s.Hospital.Name)
		}
	}
}
