# User Profile API

A production-ready RESTful API built with Go, GoFiber, and Supabase (PostgreSQL) that manages user profiles with name and date of birth. The API calculates user age dynamically in Go based on their DOB.

## Features

✅ **Full CRUD Operations** - Create, Read, Update, Delete users  
✅ **Dynamic Age Calculation** - Age calculated in Go, not stored in database  
✅ **Clean Architecture** - Separation of concerns (Handler → Service → Repository)  
✅ **Type-Safe Database Access** - Using SQLC for generated queries  
✅ **Production-Ready Logging** - Structured logging with Uber Zap  
✅ **Request Validation** - Using go-playground/validator  
✅ **Middleware Stack** - Request ID, logging, error handling, CORS  
✅ **Graceful Shutdown** - Proper cleanup on termination  
✅ **Pagination Support** - Efficient listing of users  
✅ **Health Check Endpoint** - Monitor application status  
✅ **Pagination to /users** - (API endpoint: GET /users?limit=10&offset=0)

## Prerequisites

- **Go** 1.21 or higher
- **SQLC** - Install with: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

### 2. Local Development Setup

1. **Clone the repository and Navigate to Folder**

```bash
git clone https://github.com/mrDeepakk/user-profile-api.git
cd user-profile-api
```

2. **Install dependencies**

```bash
go mod download
```

3. **Configure environment variables**

**Create `.env` file (Project Root)**

```bash
DATABASE_URL=postgresql://postgres:[PASSWORD]@db.[PROJECT-REF].supabase.co:5432/postgres  // user your own database url
PORT=3000
LOG_LEVEL=info
```

4. **Generate SQLC code**

```bash
cd db\sqlc
sqlc generate
cd ..\..
```

5. **Run the application**

```bash
go run cmd/server/main.go
```

The API will start on `http://localhost:3000`
