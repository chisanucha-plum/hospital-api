# AGNOS Hospital Middleware API

## ภาพรวมของระบบ
Hospital Middleware API สำหรับการค้นหาและจัดการข้อมูลผู้ป่วยจากระบบโรงพยาบาล (HIS - Hospital Information System)

## คุณสมบัติ
- ค้นหาข้อมูลผู้ป่วยด้วย National ID หรือ Passport ID
- จัดการข้อมูลผู้ป่วย (CRUD operations)
- รองรับข้อมูลภาษาไทยและภาษาอังกฤษ
- REST API ตามมาตรฐาน
- Database connection pooling
- Graceful shutdown

## โครงสร้างโปรเจค
```
agnos-backend/
├── cmd/
│   └── server/          # Main application
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP handlers (controllers)
│   ├── models/          # Data models และ DTOs
│   ├── repository/      # Data access layer
│   └── services/        # Business logic layer
├── database/            # Database schema และ migrations
├── go.mod
├── go.sum
└── README.md
```

## เทคโนโลยีที่ใช้
- **Go 1.24+** - Programming language
- **Gorilla Mux** - HTTP router
- **PostgreSQL** - Database
- **lib/pq** - PostgreSQL driver

## การติดตั้งและรัน

### 1. ติดตั้ง Dependencies
```bash
go mod download
```

### 2. ตั้งค่า Database
```bash
# สร้าง database
createdb agnos_hospital

# รัน schema
psql -d agnos_hospital -f database/schema.sql
```

### 3. ตั้งค่า Environment Variables
```bash
# คัดลอกไฟล์ environment
cp .env.example .env

# แก้ไขค่าต่างๆ ใน .env
```

### 4. รันโปรแกรม
```bash
go run cmd/server/main.go
```

## API Endpoints

### ค้นหาผู้ป่วย
```
GET /api/v1/patient/search/{id}
```
**Parameters:**
- `id` - National ID หรือ Passport ID

**Response:**
```json
{
  "success": true,
  "message": "Patient found",
  "data": {
    "first_name_th": "สมชาย",
    "middle_name_th": null,
    "last_name_th": "ใจดี",
    "first_name_en": "Somchai",
    "middle_name_en": null,
    "last_name_en": "Jaidee",
    "date_of_birth": "1990-05-15",
    "patient_hn": "HN001001",
    "national_id": "1234567890123",
    "passport_id": null,
    "phone_number": null,
    "email": null,
    "gender": "M"
  }
}
```

### สร้างผู้ป่วยใหม่
```
POST /api/v1/patient
```

### อัพเดทข้อมูลผู้ป่วย
```
PUT /api/v1/patient/{id}
```

### ลบข้อมูลผู้ป่วย (Soft Delete)
```
DELETE /api/v1/patient/{id}
```

### Health Check
```
GET /health
```

## การใช้งาน

### ตัวอย่างการค้นหาผู้ป่วย
```bash
curl -X GET "http://localhost:8002/api/v1/patient/search/1234567890123"
```

### ตัวอย่างการสร้างผู้ป่วยใหม่
```bash
curl -X POST "http://localhost:8002/api/v1/patient" \
  -H "Content-Type: application/json" \
  -d '{
    "national_id": "1111111111111",
    "first_name_th": "ทดสอบ",
    "last_name_th": "ระบบ",
    "first_name_en": "Test",
    "last_name_en": "System",
    "date_of_birth": "1995-01-01T00:00:00Z",
    "patient_hn": "HN002001",
    "gender": "M",
    "hospital_id": 1
  }'
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| SERVER_HOST | Server host | localhost |
| SERVER_PORT | Server port | 8002 |
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | postgres |
| DB_PASSWORD | Database password | (required) |
| DB_NAME | Database name | agnos_hospital |
| DB_SSLMODE | SSL mode | disable |

## การพัฒนา

### รันในโหมด Development
```bash
# With auto-reload (ถ้าติดตั้ง air)
air

# หรือรันปกติ
go run cmd/server/main.go
```

### การทดสอบ
```bash
go test ./...
```

### Build Production
```bash
go build -o bin/server cmd/server/main.go
```

## Security Notes
- ใช้ HTTPS ใน production
- ตั้งค่า CORS อย่างเหมาะสม
- เพิ่ม authentication/authorization ตามความต้องการ
- Validate input data อย่างเข้มงวด

## License
ลิขสิทธิ์เป็นของบริษัท AGNOS
