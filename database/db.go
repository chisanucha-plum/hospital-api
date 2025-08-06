package database

import (
	"fmt"
	"hospital-api/internal/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB() (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.Envs.DBHost,
		configs.Envs.DBPort,
		configs.Envs.DBUser,
		configs.Envs.DBPassword,
		configs.Envs.DBName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

// InitSchema สร้าง tables ถ้ายังไม่มี
func InitSchema(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS hospitals (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS staff (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		hospital_id INTEGER REFERENCES hospitals(id)
	);

	CREATE TABLE IF NOT EXISTS patients (
		id SERIAL PRIMARY KEY,
		hospital_id INTEGER REFERENCES hospitals(id),
		first_name_th VARCHAR(50),
		middle_name_th VARCHAR(50),
		last_name_th VARCHAR(50),
		first_name_en VARCHAR(50),
		middle_name_en VARCHAR(50),
		last_name_en VARCHAR(50),
		date_of_birth DATE,
		patient_hn VARCHAR(20) UNIQUE,
		national_id VARCHAR(20) UNIQUE,
		passport_id VARCHAR(20) UNIQUE,
		phone_number VARCHAR(20),
		email VARCHAR(100),
		gender CHAR(1) CHECK (gender IN ('M', 'F'))
	);`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %v", err)
	}
	return nil
}
