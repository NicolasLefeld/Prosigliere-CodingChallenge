
# Prosigliere Blog API

This is a small REST API built for a code challenge. It allows managing blog posts and comments. The API is built in Go using GORM, Chi, and PostgreSQL. It includes a simple API key authentication mechanism.


## 📦 Tech Stack

- Go 1.23.4
- GORM (ORM)
- Chi (HTTP router)
- PostgreSQL
- Docker & Docker Compose


## 🚀 How to Run the Project

### 🔧 Option 1: Run Locally with `.env`

1. Create a `.env` file in the root of the project:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=blogdb
PORT=8080
API_KEY=123456
```

2. Make sure PostgreSQL is running and the `blogdb` database exists.

3. Run the project:

```bash
go mod tidy
go run main.go
```

---

### 🐳 Option 2: Run with Docker Compose

This will spin up the API and PostgreSQL in one step.

```bash
docker compose up --build
```

The API will be available at:  
http://localhost:8080


## 🔐 Authentication

All endpoints require an API key via the following header:

```
X-API-Key: 123456
```

You can change this value in `.env` or `docker-compose.yml`.


## 📬 Postman

You can use the provided Postman collection to test the API:  
**`prosigliere-blog.postman_collection.json`**

### Endpoints available:

- `GET /api/posts` – List all posts with comment count
- `POST /api/posts` – Create a new post
- `GET /api/posts/{id}` – Retrieve a post by ID, including comments
- `POST /api/posts/{id}/comments` – Add a comment to a post

All requests include the `X-API-Key` header.


## 📂 Project Structure

```
.
├── main.go                # Entry point
├── config/                # DB connection and migration
├── models/                # GORM models
├── handlers/              # HTTP handlers for posts and comments
├── routes/                # Chi router setup
├── middleware/            # API key auth middleware
├── Dockerfile             # Docker build for Go app
├── docker-compose.yml     # Runs app + PostgreSQL
├── .env                   # Optional, local config
└── prosigliere-blog.postman_collection.json
```


## ✅ Ready to Go

This project is ready to run with Docker or locally. It auto-creates tables and includes a Postman collection for testing.
