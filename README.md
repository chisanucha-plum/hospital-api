# Hospital API Management System

## ภาพรวมของระบบ
Hospital API สำหรับการจัดการข้อมูลโรงพยาบาล เจ้าหน้าที่ และผู้ป่วย พร้อมระบบ Authentication และ Authorization ที่ครบถ้วน พัฒนาด้วย Go และ PostgreSQL ใช้ architecture แบบ Clean Code และ Type-Safe Models

## คุณสมบัติหลัก
- 🏥 **Hospital Management** - จัดการข้อมูลโรงพยาบาลพร้อม relationship
- 👨‍⚕️ **Staff Authentication** - ระบบสมัครสมาชิกและเข้าสู่ระบบด้วย JWT
- 🔍 **Advanced Patient Search** - ค้นหาผู้ป่วยแบบ multi-field ด้วย query parameters
- 🛡️ **Type-Safe JWT Claims** - JWT authentication ด้วย structured claims
- 🌍 **Multilingual Support** - รองรับข้อมูลภาษาไทยและภาษาอังกฤษ
- 📊 **GORM ORM** - Auto migration และ type-safe database operations
- 🔗 **Foreign Key Relationships** - Hospital ↔ Staff ↔ Patient
- 🐳 **Docker Ready** - Docker Compose สำหรับ development และ production
- 📝 **Structured API Models** - Type-safe request/response models

## เทคโนโลยีที่ใช้
- **Go 1.22+** - Core programming language
- **Gin Framework** - HTTP router พร้อม middleware และ route grouping
- **GORM** - ORM สำหรับ database operations
- **PostgreSQL** - Primary database with foreign key constraints
- **JWT (golang-jwt/jwt/v5)** - Structured authentication claims
- **bcrypt** - Secure password hashing
- **Docker & Docker Compose** - Containerization และ development environment

## โครงสร้างโปรเจค (Clean Architecture)
```
hospital-api/
├── cmd/
│   └── main.go                    # Application entry point
├── internal/
│   ├── configs/
│   │   └── envs.go               # Environment configuration
│   ├── handlers/
│   │   ├── router.go             # Route grouping & middleware setup
│   │   ├── staff.go              # Staff endpoints (create, login)
│   │   └── patient.go            # Patient endpoints (search)
│   ├── middleware/
│   │   └── auth.go               # JWT authentication middleware
│   ├── models/
│   │   ├── api.go                # Structured API request/response models
│   │   ├── hospital.go           # Hospital domain model
│   │   └── user.go               # Staff & Patient domain models
│   └── services/
│       ├── auth.go               # JWT generation & validation
│       ├── staff.go              # Staff business logic
│       └── painet.go             # Patient business logic
├── database/
│   ├── db.go                     # GORM connection & auto migration
│   └── seed.go                   # Mock data seeding
├── nginx/
│   └── nginx.conf                # Nginx reverse proxy configuration
├── Dockerfile                    # Multi-stage Docker build
├── compose.yaml                  # Docker Compose setup
├── .env.example                  # Environment template
└── README.md
```

## การติดตั้งและรัน

### 1. Clone Repository
```bash
git clone https://github.com/chisanucha-plum/hospital-api.git
cd hospital-api
```

### 2. ตั้งค่า Environment Variables
```bash
# คัดลอกไฟล์ environment template
cp .env.example .env

# แก้ไขค่าต่างๆ ใน .env ตามต้องการ
# JWT_SECRET=your-secret-key-here
# DB_HOST=postgres_server
# DB_USER=postgres
# DB_PASSWORD=password
# DB_NAME=hospital_db
```

### 3. รันด้วย Docker Compose (แนะนำ)
```bash
# Build และ start ทุก service
docker-compose up --build

# หรือรันใน background
docker-compose up -d --build

# ดู logs
docker-compose logs -f

# หยุด service
docker-compose down
```

### 4. รันแบบ Local Development
```bash
# ติดตั้ง dependencies
go mod download

# ตั้งค่า PostgreSQL local และแก้ไข .env
# DB_HOST=localhost

# รันโปรแกรม
go run cmd/main.go
```

## API Endpoints

### 🔐 Staff Authentication

#### สมัครสมาชิกเจ้าหน้าที่ (ใช้ hospital_name)
```http
POST /api/v1/staff/create
Content-Type: application/json

{
  "username": "admin123",
  "password": "securepassword",
  "hospital_name": "โรงพยาบาลศิริราช"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Staff created successfully",
  "data": {
    "message": "Staff created successfully",
    "staff_id": 1
  }
}
```

#### เข้าสู่ระบบ
```http
POST /api/v1/staff/login
Content-Type: application/json

{
  "username": "admin123",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 🏥 Patient Management (ต้องใช้ JWT Token)

#### ค้นหาผู้ป่วยด้วย ID (National ID หรือ Passport)
```http
GET /api/v1/patients/search/{id}
Authorization: Bearer {JWT_TOKEN}
```

**ตัวอย่าง:**
```bash
GET /api/v1/patients/search/1234567890123
```

**Response:**
```json
{
  "success": true,
  "message": "Patient found",
  "data": {
    "national_id": "1234567890123",
    "patient_hn": "HN0001",
    "first_name_th": "สมชาย",
    "last_name_th": "ใจดี",
    "first_name_en": "Somchai",
    "last_name_en": "Jaidee",
    "date_of_birth": "1990-01-01T00:00:00Z",
    "passport_id": "",
    "phone_number": "0812345678",
    "email": "somchai@example.com",
    "gender": "M",
    "hospital_id": 1
  }
}
```

#### ค้นหาผู้ป่วยแบบ Multi-field (Query Parameters)
```http
GET /api/v1/patients/search?field=value&field2=value2
Authorization: Bearer {JWT_TOKEN}
```

**ตัวอย่างการใช้งาน:**
```bash
# ค้นหาด้วยชื่อ (partial match)
GET /api/v1/patients/search?first_name_th=สมชาย

# ค้นหาด้วยเบอร์โทร (exact match)
GET /api/v1/patients/search?phone_number=0812345678

# ค้นหาด้วยหลายเงื่อนไข
GET /api/v1/patients/search?first_name_th=สม&last_name_th=ใจ&email=som

# ค้นหาด้วย HN
GET /api/v1/patients/search?patient_hn=HN0001
```

**Response:**
```json
{
  "success": true,
  "message": "Patients found",
  "data": {
    "patients": [
      {
        "national_id": "1234567890123",
        "patient_hn": "HN0001",
        "first_name_th": "สมชาย",
        "last_name_th": "ใจดี"
        // ... other fields
      }
    ],
    "count": 1
  }
}
```

### Search Fields ที่รองรับ

| Field | Type | Match Type | Description |
|-------|------|------------|-------------|
| `national_id` | string | Exact | เลขบัตรประชาชน |
| `passport_id` | string | Exact | เลขหนังสือเดินทาง |
| `patient_hn` | string | Exact | Hospital Number |
| `phone_number` | string | Exact | เบอร์โทรศัพท์ |
| `first_name_th` | string | Partial | ชื่อภาษาไทย |
| `last_name_th` | string | Partial | นามสกุลภาษาไทย |
| `first_name_en` | string | Partial | ชื่อภาษาอังกฤษ |
| `last_name_en` | string | Partial | นามสกุลภาษาอังกฤษ |
| `email` | string | Partial | อีเมล |

## การใช้งานจริง

### Mock Data ที่มีในระบบ

เมื่อรันระบบครั้งแรก จะมีการสร้าง Mock Data อัตโนมัติ:

#### 🏥 **Hospitals (5 โรงพยาบาล)**
1. โรงพยาบาลศิริราช (ID: 1)
2. โรงพยาบาลจุฬาลงกรณ์ (ID: 2) 
3. โรงพยาบาลรามาธิบดี (ID: 3)
4. โรงพยาบาลเวชศาสตร์เขตร้อน (ID: 4)
5. Bangkok Hospital (ID: 5)

#### 👨‍⚕️ **Staff Accounts (6 บัญชี)**
| Username | Password | Hospital |
|----------|----------|----------|
| `admin_siriraj` | `password123` | โรงพยาบาลศิริราช |
| `admin_chula` | `password123` | โรงพยาบาลจุฬาลงกรณ์ |
| `admin_rama` | `password123` | โรงพยาบาลรามาธิบดี |
| `doctor_tropical` | `password123` | โรงพยาบาลเวชศาสตร์เขตร้อน |
| `nurse_bangkok` | `password123` | Bangkok Hospital |
| `admin_siriraj2` | `password123` | โรงพยาบาลศิริราช |

#### 🏥 **Patients (10 ผู้ป่วย)**
| National ID | Patient HN | Name (TH) | Name (EN) | Hospital |
|-------------|------------|-----------|-----------|----------|
| `1234567890123` | `HN0001` | สมชาย ใจดี | Somchai Jaidee | ศิริราช |
| `2345678901234` | `HN0002` | สมหญิง สวยงาม | Somying Suaynam | ศิริราช |
| `3456789012345` | `HN0003` | วิชัย เก่งมาก | Wichai Kengmak | จุฬาลงกรณ์ |
| `4567890123456` | `HN0004` | วันทนา รักดี | Wantana Rakdee | จุฬาลงกรณ์ |
| `5678901234567` | `HN0005` | ประพจน์ สมประสงค์ | Prapot Somprasong | รามาธิบดี |
| และอื่นๆ... | | | | |

### ตัวอย่างการทำงานแบบ Step-by-Step

#### 1. สมัครสมาชิกเจ้าหน้าที่ (หรือใช้ Mock Account)
```bash
# วิธีที่ 1: สมัครใหม่
curl -X POST "http://localhost:8002/api/v1/staff/create" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin123",
    "password": "mypassword",
    "hospital_name": "โรงพยาบาลศิริราช"
  }'

# วิธีที่ 2: ใช้ Mock Account (แนะนำ)
# Username: admin_siriraj, Password: password123
```

#### 2. เข้าสู่ระบบเพื่อรับ JWT Token
```bash
# ใช้ Mock Account
curl -X POST "http://localhost:8002/api/v1/staff/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin_siriraj",
    "password": "password123"
  }'
```

#### 3. ค้นหาผู้ป่วยด้วย National ID (ใช้ Mock Data)
```bash
# ค้นหาผู้ป่วย สมชาย ใจดี
curl -X GET "http://localhost:8002/api/v1/patients/search/1234567890123" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

#### 4. ค้นหาผู้ป่วยด้วยชื่อ (Query Parameters)
```bash
# ค้นหาด้วยชื่อภาษาไทย
curl -X GET "http://localhost:8002/api/v1/patients/search?first_name_th=สมชาย" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"

# ค้นหาด้วย Hospital Number
curl -X GET "http://localhost:8002/api/v1/patients/search?patient_hn=HN0001" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | API server port | 8002 | No |
| `DB_HOST` | Database host | postgres_server | Yes |
| `DB_PORT` | Database port | 5432 | Yes |
| `DB_USER` | Database username | postgres | Yes |
| `DB_PASSWORD` | Database password | - | Yes |
| `DB_NAME` | Database name | hospital_db | Yes |
| `JWT_SECRET` | JWT signing secret (ต้องแข็งแรง) | - | Yes |

## Database Schema

### ERD (Entity Relationship Diagram)
```
┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
│    Hospital     │       │      Staff      │       │     Patient     │
├─────────────────┤   ┌──→├─────────────────┤   ┌──→├─────────────────┤
│ id (PK)         │   │   │ id (PK)         │   │   │ national_id(PK) │
│ name (unique)   │   │   │ username        │   │   │ patient_hn      │
│ created_at      │   │   │ password (hash) │   │   │ first_name_th   │
│ updated_at      │   │   │ hospital_id(FK) │───┘   │ last_name_th    │
└─────────────────┘   │   │ created_at      │       │ first_name_en   │
                      │   │ updated_at      │       │ last_name_en    │
                      │   └─────────────────┘       │ date_of_birth   │
                      │                             │ passport_id     │
                      │                             │ phone_number    │
                      │                             │ email           │
                      │                             │ gender (M/F)    │
                      │                             │ hospital_id(FK) │───┘
                      │                             │ created_at      │
                      │                             │ updated_at      │
                      │                             └─────────────────┘
                      └─────────────────────────────────────────────────
```

### Primary Keys และ Constraints
- **Hospital**: `id` (auto-increment), `name` (unique)
- **Staff**: `id` (auto-increment), `username` + `hospital_id` (unique)
- **Patient**: `national_id` (string PK), `patient_hn` + `hospital_id` (unique)
- **Gender**: Enum constraint (M, F เท่านั้น)
- **Foreign Keys**: Staff.hospital_id → Hospital.id, Patient.hospital_id → Hospital.id

## Type-Safe Models

### API Models (Structured)
```go
// Request Models
type CreateStaffRequest struct {
    Username     string `json:"username" binding:"required"`
    Password     string `json:"password" binding:"required"`
    HospitalName string `json:"hospital_name" binding:"required"`
}

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// Response Models
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// JWT Claims (Type-Safe)
type JWTClaims struct {
    StaffID    int `json:"staff_id"`
    HospitalID int `json:"hospital_id"`
    jwt.RegisteredClaims
}
```

## Architecture & Design Patterns

### Clean Architecture Layers
- **Handlers**: HTTP request/response handling
- **Services**: Business logic และ data validation
- **Models**: Domain models และ API models
- **Database**: GORM ORM และ PostgreSQL

### Design Patterns ที่ใช้
- **Dependency Injection**: ใช้ `Depends()` pattern
- **Repository Pattern**: Service layer abstracts database operations
- **Factory Pattern**: NewService constructors
- **Middleware Pattern**: JWT authentication
- **Route Grouping**: API versioning และ middleware organization

## การพัฒนา

### Local Development
```bash
# ด้วย Docker Compose (แนะนำ)
docker-compose up --build

# Live reload ด้วย air (ถ้าติดตั้งแล้ว)
air

# รัน tests
go test ./...

# Format code
gofmt -w .
```

### Testing
```bash
# รัน unit tests ทั้งหมด
go test ./...

# Test เฉพาะ package
go test ./internal/services/
go test ./internal/handlers/

# Test พร้อม coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

### Build Production
```bash
# Build binary
go build -o bin/hospital-api cmd/main.go

# Build Docker image
docker build -t hospital-api:latest .
```

## Docker Deployment

### Development Environment
```bash
# Start all services
docker-compose up --build

# View logs
docker-compose logs -f api

# Reset database
docker-compose down -v && docker-compose up --build
```

### Production Environment
```bash
# Build production image
docker build -t hospital-api:production .

# Run with production env
docker run -d \
  --name hospital-api-prod \
  -p 8002:8002 \
  --env-file .env.production \
  hospital-api:production
```

## Security & Best Practices

### 🔐 Security Features
- **Structured JWT Claims** - Type-safe authentication tokens
- **bcrypt Password Hashing** - Industry-standard password security
- **Foreign Key Constraints** - Data integrity enforcement
- **Input Validation** - Request model validation
- **SQL Injection Protection** - Parameterized queries ผ่าน GORM
- **Hospital Isolation** - Data segregation ตาม hospital_id

### 📋 Production Checklist
- [ ] ใช้ HTTPS/TLS ใน production
- [ ] ตั้งค่า CORS policy อย่างเหมาะสม
- [ ] เปลี่ยน JWT_SECRET เป็นค่าที่แข็งแรง (256-bit)
- [ ] ตั้งค่า rate limiting และ request throttling
- [ ] เพิ่ม structured logging และ monitoring
- [ ] Database backup automation
- [ ] Health check endpoints
- [ ] Error tracking และ alerting

## Performance Optimization

### Database Indexing
```sql
-- สำหรับ patient search performance
CREATE INDEX idx_patient_hospital_id ON user_patients(hospital_id);
CREATE INDEX idx_patient_names ON user_patients(first_name_th, last_name_th);
CREATE INDEX idx_patient_phone ON user_patients(phone_number);
CREATE INDEX idx_patient_email ON user_patients(email);
```

### Caching Strategy
- **JWT Token Caching**: In-memory token validation
- **Database Connection Pooling**: GORM connection pool
- **Query Result Caching**: Redis (แนะนำสำหรับ production)

### Monitoring Metrics
- Response time per endpoint
- Database query performance
- JWT token validation time
- Memory และ CPU usage
- Error rate และ status codes

## Contributing

### Code Style Guidelines
- ใช้ `gofmt` สำหรับ formatting
- ตั้งชื่อตัวแปรและฟังก์ชันแบบ camelCase
- เขียน godoc comments สำหรับ public functions
- ใช้ structured logging
- เขียน unit tests สำหรับฟังก์ชันใหม่

### Pull Request Process
1. Fork repository นี้
2. สร้าง feature branch (`git checkout -b feature/amazing-feature`)
3. Commit การเปลี่ยนแปลง (`git commit -m 'Add amazing feature'`)
4. Push ไปยัง branch (`git push origin feature/amazing-feature`)
5. เปิด Pull Request พร้อม description ที่ชัดเจน

## Troubleshooting

### ปัญหาที่พบบ่อย

#### 1. "Connection refused" - ไม่สามารถเชื่อมต่อ database
```bash
# เช็คว่า PostgreSQL รันอยู่หรือไม่
docker-compose ps

# restart database service
docker-compose restart postgres_server

# ดู logs ของ database
docker-compose logs postgres_server
```

#### 2. "Foreign key constraint" error
```bash
# ลบข้อมูลเดิมและ migrate ใหม่
docker-compose down -v
docker-compose up --build

# หรือ reset เฉพาะ database
docker-compose exec postgres_server psql -U postgres -d hospital_db -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
```

#### 3. "JWT token invalid"
- ตรวจสอบว่า JWT_SECRET ใน .env ตรงกันหรือไม่
- ตรวจสอบว่า token ยังไม่หมดอายุ (24 ชั่วโมง)
- ตรวจสอบ Authorization header format: `Bearer <token>`

#### 4. "Hospital not found" error
```bash
# เช็คว่ามี hospital ในระบบหรือไม่
docker-compose exec postgres_server psql -U postgres -d hospital_db -c "SELECT * FROM hospitals;"

# หรือใส่ mock data
docker-compose exec api go run database/seed.go
```

## Version History

### v2.0.0 (Current)
- ✅ Type-safe JWT claims structure
- ✅ Structured API request/response models
- ✅ Query parameters สำหรับ patient search
- ✅ Route grouping และ API versioning
- ✅ Enhanced error handling

### v1.0.0 (Previous)
- Basic CRUD operations
- JWT authentication
- Docker containerization
- GORM integration

---

## License
Copyright © 2025 Hospital API Management System  
All rights reserved.

## Contact & Support
สำหรับข้อสงสัยหรือปัญหาการใช้งาน:
- **GitHub Issues**: [Create an issue](https://github.com/chisanucha-plum/hospital-api/issues)
- **Documentation**: [API Documentation](https://github.com/chisanucha-plum/hospital-api/wiki)

---
**Made with ❤️ for Thai Healthcare System**  
*Powered by Go, PostgreSQL, and Docker*

## การติดตั้งและรัน

### 1. Clone Repository
```bash
git clone https://github.com/chisanucha-plum/hospital-api.git
cd hospital-api
```

### 2. ตั้งค่า Environment Variables
```bash
# คัดลอกไฟล์ environment template
cp .env.example .env

# แก้ไขค่าต่างๆ ใน .env ตามต้องการ
```

### 3. รันด้วย Docker Compose (แนะนำ)
```bash
# Build และ start ทุก service
docker-compose up --build

# หรือรันใน background
docker-compose up -d --build

# หยุด service
docker-compose down
```

### 4. รันแบบ Local Development (ถ้าไม่ใช้ Docker)
```bash
# ติดตั้ง dependencies
go mod download

# ตั้งค่า PostgreSQL local และแก้ไข .env
# DB_HOST=localhost
# DB_PORT=5432

# รันโปรแกรม
go run cmd/main.go
```

## API Endpoints

### 🔐 Staff Authentication

#### สมัครสมาชิกเจ้าหน้าที่
```
POST /staff/create
Content-Type: application/json

{
  "username": "admin123",
  "password": "securepassword",
  "hospital_id": 1
}
```

#### เข้าสู่ระบบ
```
POST /staff/login
Content-Type: application/json

{
  "username": "admin123",
  "password": "securepassword"
}
```
**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 🏥 Patient Management (ต้องใช้ JWT)

#### ค้นหาผู้ป่วยด้วย ID
```
GET /patient/search/{national_id}
Authorization: Bearer {JWT_TOKEN}
```
**Response:**
```json
{
  "national_id": "1234567890123",
  "patient_hn": "HN0001",
  "first_name_th": "สมชาย",
  "middle_name_th": "",
  "last_name_th": "ใจดี",
  "first_name_en": "Somchai",
  "middle_name_en": "",
  "last_name_en": "Jaidee",
  "date_of_birth": "1990-01-01T00:00:00Z",
  "passport_id": "",
  "phone_number": "0812345678",
  "email": "somchai@example.com",
  "gender": "M",
  "hospital_id": 1
}
```

#### ค้นหาผู้ป่วยแบบ Multi-field (OR Search)
```
POST /patient/search
Authorization: Bearer {JWT_TOKEN}
Content-Type: application/json

{
  "first_name_th": "สมชาย",
  "phone_number": "0812345678",
  "patient_hn": "HN0001"
}
```
**Response:**
```json
[
  {
    "national_id": "1234567890123",
    "patient_hn": "HN0001",
    "first_name_th": "สมชาย",
    "last_name_th": "ใจดี",
    // ... other fields
  }
]
```

#### ค้นหาผู้ป่วยแบบ JSON Body (รองรับทั้ง GET และ POST)
```
GET /patient/search
Authorization: Bearer {JWT_TOKEN}
Content-Type: application/json

{
  "first_name_th": "สมชาย",
  "last_name_th": "ใจดี"
}
```
**Response:**
```json
{
  "success": true,
  "message": "Patients found",
  "data": {
    "count": 1,
    "patients": [
      {
        "national_id": "1234567890123",
        "patient_hn": "HN0001",
        "first_name_th": "สมชาย",
        "last_name_th": "ใจดี",
        "first_name_en": "Somchai",
        "last_name_en": "Jaidee",
        "date_of_birth": "1990-01-01T00:00:00+07:00",
        "gender": "M",
        "hospital_id": "1"
      }
    ]
  }
}
```

## การใช้งาน

### ตัวอย่างการเข้าสู่ระบบและค้นหาผู้ป่วย

#### 1. สมัครสมาชิกเจ้าหน้าที่
```bash
curl -X POST "http://localhost:8081/staff/create" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin123",
    "password": "mypassword",
    "hospital_id": 1
  }'
```

#### 2. เข้าสู่ระบบเพื่อรับ JWT Token
```bash
curl -X POST "http://localhost:8081/staff/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin123",
    "password": "mypassword"
  }'
```

#### 3. ค้นหาผู้ป่วยด้วย National ID
```bash
curl -X GET "http://localhost:8081/patient/search/1234567890123" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

#### 4. ค้นหาผู้ป่วยแบบ Multi-field
```bash
curl -X POST "http://localhost:8081/patient/search" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "first_name_th": "สมชาย"
  }'
```

#### 5. ค้นหาผู้ป่วยด้วย JSON Body (GET)
```bash
curl -X GET "http://localhost:8002/patient/search" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "first_name_th": "สมชาย",
    "last_name_th": "ใจดี"
  }'
```

## Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | API server port | 8002 | No |
| `DB_HOST` | Database host | localhost | Yes |
| `DB_PORT` | Database port | 5432 | Yes |
| `DB_USER` | Database username | - | Yes |
| `DB_PASSWORD` | Database password | - | Yes |
| `DB_NAME` | Database name | - | Yes |
| `JWT_SECRET` | JWT signing secret | - | Yes |

## Database Schema

### ERD (Entity Relationship Diagram)
```
┌─────────────┐       ┌─────────────┐       ┌─────────────┐
│   Hospital  │       │    Staff    │       │   Patient   │
├─────────────┤   ┌──→├─────────────┤   ┌──→├─────────────┤
│ id (PK)     │   │   │ id (PK)     │   │   │ national_id │
│ name        │   │   │ username    │   │   │ patient_hn  │
│ created_at  │   │   │ password    │   │   │ first_name  │
│ updated_at  │   │   │ hospital_id │───┘   │ last_name   │
└─────────────┘   │   │ created_at  │       │ hospital_id │───┘
                  │   │ updated_at  │       │ gender      │
                  │   └─────────────┘       │ created_at  │
                  │                         │ updated_at  │
                  └─────────────────────────└─────────────┘
```

### หลักการ Relationships
- `Hospital` → `Staff` (One-to-Many)
- `Hospital` → `Patient` (One-to-Many)  
- `Staff.hospital_id` และ `Patient.hospital_id` เป็น Foreign Keys

## การพัฒนา

### รันในโหมด Development
```bash
# ด้วย Docker Compose (แนะนำ)
docker-compose up --build

# หรือ local development
go run cmd/main.go

# ถ้าติดตั้ง air สำหรับ auto-reload
air
```

### การทดสอบ
```bash
# รัน unit tests ทั้งหมด
go test ./...

# รัน test เฉพาะ package
go test ./internal/services/
go test ./internal/router/

# รัน test พร้อม coverage
go test -cover ./...
```

### Build Production
```bash
# Build binary
go build -o bin/hospital-api cmd/main.go

# หรือ build Docker image
docker build -t hospital-api .
```

## Docker Deployment

### Development
```bash
docker-compose up --build
```

### Production
```bash
# Build production image
docker build -t hospital-api:latest .

# Run with custom env
docker run -d \
  --name hospital-api \
  -p 8002:8002 \
  --env-file .env \
  hospital-api:latest
```

## Security & Best Practices

### 🔐 Security Features
- **JWT Authentication** - ป้องกันการเข้าถึงข้อมูลโดยไม่ได้รับอนุญาต
- **bcrypt Password Hashing** - เข้ารหัสรหัสผ่านอย่างปลอดภัย
- **Foreign Key Constraints** - รักษาความสมบูรณ์ของข้อมูล
- **Input Validation** - ตรวจสอบข้อมูลที่เข้ามาทุกรายการ

### 📋 Production Checklist
- [ ] ใช้ HTTPS/TLS ใน production
- [ ] ตั้งค่า CORS อย่างเหมาะสม
- [ ] เปลี่ยน JWT_SECRET เป็นค่าที่แข็งแรง
- [ ] ตั้งค่า rate limiting
- [ ] เพิ่ม logging และ monitoring
- [ ] Backup database เป็นประจำ

## Contributing

### Code Style
- ใช้ `gofmt` สำหรับ formatting
- ตั้งชื่อตัวแปรและฟังก์ชันแบบ camelCase
- เขียน unit tests สำหรับฟังก์ชันใหม่
- เขียน comment ภาษาไทยสำหรับ business logic

### Pull Request Process
1. Fork repository นี้
2. สร้าง feature branch (`git checkout -b feature/amazing-feature`)
3. Commit การเปลี่ยนแปลง (`git commit -m 'Add amazing feature'`)
4. Push ไปยัง branch (`git push origin feature/amazing-feature`)
5. เปิด Pull Request

## Troubleshooting

### ปัญหาที่พบบ่อย

#### 1. "Connection refused" - ไม่สามารถเชื่อมต่อ database
```bash
# เช็คว่า PostgreSQL รันอยู่หรือไม่
docker-compose ps

# restart database service
docker-compose restart postgres_server
```

#### 2. "Foreign key constraint" error
```bash
# ลบข้อมูลเดิมและ migrate ใหม่
docker-compose down -v
docker-compose up --build
```

#### 3. "JWT token invalid"
- ตรวจสอบว่า JWT_SECRET ใน .env ตรงกันหรือไม่
- ตรวจสอบว่า token ยังไม่หมดอายุ (24 ชั่วโมง)

## Performance & Monitoring

### Metrics ที่ควรติดตาม
- Response time per endpoint
- Database connection pool usage
- Memory และ CPU usage
- Error rate และ logs

### Scaling Recommendations
- ใช้ Database connection pooling
- เพิ่ม Redis สำหรับ caching
- Load balancer สำหรับ multiple instances
- Database indexing สำหรับ search queries

## License
Copyright © 2025 Hospital API Management System  
All rights reserved.

## Contact
สำหรับข้อสงสัยหรือปัญหาการใช้งาน:
- GitHub Issues: [Create an issue](https://github.com/chisanucha-plum/hospital-api/issues)
- Email: support@hospital-api.com

---
**Made with ❤️ for Thai Healthcare System**
