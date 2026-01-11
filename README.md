# REKA Storage Service

A robust file storage service built with Go, utilizing MinIO for object storage and MongoDB for metadata management.

## 🚀 Technology Stack

- **Language:** Go 1.22
- **Framework:** Gin Web Framework
- **Database:** MongoDB
- **Object Storage:** MinIO
- **Authentication:** JWT (JSON Web Tokens).
- **Configuration:** Godotenv 

## 📋 Prerequisites 

Before running the application, ensure you have the following installed:

- [Go](https://go.dev/dl/) (version 1.22 or later)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [MinIO](https://min.io/download)

## 🛠️ Installation & Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd reka-storage
   ```

2. **Install Dependencies**
   ```bash
   go mod download
   ```

3. **Environment Configuration**
   Create a `.env` file in the root directory based on the following template:

   ```env
   # App Configuration
   APP_PORT=4000

   # MongoDB Configuration
   MONGO_URI=mongodb://localhost:27017
   MONGO_DB=storage_db

   # MinIO Configuration
   REKASTORAGE_HOST=localhost
   REKASTORAGE_PORT=9000
   REKASTORAGE_USE_SSL=false
   REKASTORAGE_ACCESS_KEY=your_minio_access_key
   REKASTORAGE_SECRET_KEY=your_minio_secret_key
   REKASTORAGE_BUCKET=your_bucket_name
   ```

4. **Run the Application**
   ```bash
   go run cmd/server/main.go
   ```
   The server will start on port `4000` (or the port specified in `.env`).

## 🔌 API Endpoints

### Health Check
- `GET /ping` - Check if server is running.

### Authentication
- `POST /api/auth/login` - User login to obtain JWT token.

### Storage
**Note:** These endpoints require a valid JWT token in the `Authorization` header (`Bearer <token>`).

- `POST /api/storage/upload` - Upload a file to storage.

## 📁 Project Structure

```
reka-storage/
├── cmd/
│   └── server/         # Entry point of the application
├── internal/
│   ├── auth/           # Authentication logic (Handler, Service, Repo)
│   ├── storage/        # Storage logic (Handler, Service, Repo)
│   └── shared/         # Shared utilities and middleware
├── pkg/                # Public library code
├── .env                # Environment variables
├── go.mod              # Go module definition
└── README.md           # Project documentation
```
