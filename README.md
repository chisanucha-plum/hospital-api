# Hospital API Management System

## Stack
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
│   │   ├── staff.go              # Staff endpoints (create, login)
│   │   └── patient.go            # Patient endpoints (search)
│   ├── middleware/
│   │   └── auth.go               # JWT authentication middleware
│   ├── models/
│   │   ├── api.go                # Structured API request/response models
│   │   ├── hospital.go           # Hospital domain model
│   │   └── user.go               # Staff & Patient domain models
│   ├── router/
│   │   └── router.go             # Route grouping & middleware setup
│   └── services/
│       ├── auth.go               # JWT generation & validation
│       ├── staff.go              # Staff business logic
│       └── painet.go             # Patient business logic
├── database/
│   ├── db.go                     # GORM connection & auto migration
│   └── mock_data.go              # Mock data seeding
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

#### ค้นหาผปู้่วยแบบ JSON Body
```
GET /patient/search
Authorization: Bearer {JWT_TOKEN}
Content-Type: application/json
{
 "first_name_th": "สมชาย",
 "last_name_th": "ใจด"ี
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
 "last_name_th": "ใจด"ี,
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

เมื่อรันระบบครั้งแรก จะมีการสร้าง Mock Data อัตโนมัติ (แบบย่อ):

#### 🏥 **Hospitals (3 โรงพยาบาล)**
1. โรงพยาบาลศิริราช (ID: 1)
2. โรงพยาบาลจุฬาลงกรณ์ (ID: 2) 
3. โรงพยาบาลรามาธิบดี (ID: 3)

#### 👨‍⚕️ **Staff Accounts (3 บัญชี)**
| Username | Password | Hospital |
|----------|----------|----------|
| `admin1` | `password123` | โรงพยาบาลศิริราช |
| `admin2` | `password123` | โรงพยาบาลจุฬาลงกรณ์ |
| `admin3` | `password123` | โรงพยาบาลรามาธิบดี |

#### 🏥 **Patients (6 ผู้ป่วย)**
| National ID | Patient HN | Name (TH) | Name (EN) | Hospital |
|-------------|------------|-----------|-----------|----------|
| `1234567890123` | `HN001` | สมชาย ใจดี | Somchai Jaidee | ศิริราช |
| `2345678901234` | `HN002` | สมหญิง สวยงาม | Somying Suaynam | จุฬาลงกรณ์ |
| `3456789012345` | `HN003` | วิชัย เก่งมาก | Wichai Kengmak | รามาธิบดี |
| `4567890123456` | `HN004` | วันทนา รักดี | Wantana Rakdee | ศิริราช |
| `5678901234567` | `HN005` | ประพจน์ สมประสงค์ | Prapot Somprasong | จุฬาลงกรณ์ |
| `6789012345678` | `HN006` | สุภาพร เรียนดี | Supaporn Riandee | รามาธิบดี |

