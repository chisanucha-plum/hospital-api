# Hospital API Management System

## ภาพรวมของระบบ
Hospital API สำหรับการจัดการข้อมูลโรงพยาบาล, เจ้าหน้าที่, และผู้ป่วย พร้อมระบบ Authentication และ Authorization ที่ครบถ้วน

## คุณสมบัติ
- 🏥 จัดการข้อมูลโรงพยาบาล (Hospital Management)
- 👨‍⚕️ ระบบสมัครสมาชิกและเข้าสู่ระบบของเจ้าหน้าที่ (Staff Authentication)
- 🔍 ค้นหาข้อมูลผู้ป่วยแบบ Multi-field Search (National ID, Passport, Phone, etc.)
- 🛡️ JWT Authentication & Authorization
- 🌍 รองรับข้อมูลภาษาไทยและภาษาอังกฤษ
- 🔗 Foreign Key Relationships ระหว่าง Hospital-Staff-Patient
- 📊 GORM ORM พร้อม Auto Migration
- 🐳 Docker Compose สำหรับ Development และ Production

## โครงสร้างโปรเจค
```
hospital-api/
├── cmd/
│   └── main.go                    # Entry point
├── internal/
│   ├── configs/
│   │   └── envs.go               # Environment configurations
│   ├── handlers/
│   │   ├── router.go             # Route grouping & middleware setup
│   │   ├── staff.go              # Staff endpoints (create, login)
│   │   └── patient.go            # Patient endpoints (search)
│   ├── middleware/
│   │   └── auth.go               # JWT authentication middleware
│   ├── models/
│   │   ├── hospital.go           # Hospital model with relationships
│   │   └── user.go               # Staff & Patient models
│   └── services/
│       ├── auth.go               # JWT generation & validation
│       ├── staff.go              # Staff business logic
│       └── painet.go             # Patient business logic
├── database/
│   ├── db.go                     # GORM connection & auto migration
│   ├── seed.go                   # Mock data seeding
│   └── insert_mock_query.go      # SQL mock data insertion
├── nginx/
│   └── nginx.conf                # Nginx reverse proxy config
├── Dockerfile                    # Multi-stage Docker build
├── compose.yaml                  # Docker Compose setup
├── .env.example                  # Environment template
└── README.md
```

## เทคโนโลยีที่ใช้
- **Go 1.22+** - Programming language
- **Gin Framework** - HTTP router และ middleware
- **GORM** - ORM สำหรับ database operations
- **PostgreSQL** - หลักฐาน database
- **JWT (golang-jwt/jwt/v5)** - Authentication
- **bcrypt** - Password hashing
- **Docker & Docker Compose** - Containerization

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
