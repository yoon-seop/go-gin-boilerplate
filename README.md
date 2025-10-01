# Go Gin Boilerplate

A simple boilerplate project using **Go + Gin + Gorm**.  
It uses MySQL and can be easily started with Docker Compose.  
The API server can be launched via the provided `run-server.sh` script.

---

## Requirements
- Go 1.24.7
- Docker, Docker Compose
- Make (optional)

---

## 1. Start MySQL with Docker

Use the included `docker-compose.yml` file to start MySQL:

    docker-compose up -d

**Default connection info:**
- Host: `127.0.0.1`
- Port: `3306`
- User: `youruser`
- Password: `yourpassword`
- Database: `yourdb`

> MySQL data is persisted in the `mysql-data` volume.

---

## 2. Configuration

Configuration files are located under the `config/` directory:  
- `development.yaml`  
- `staging.yaml`  
- `production.yaml`  

Update database connection settings as needed.

---

## 3. Run the Server

Start the API server with the helper script:

    ./run-server.sh

Or run it directly with Go:

    go run main.go

---

## 4. Test the API

Sample HTTP requests are available in the `test.http` file.  
If you are using VSCode, you can run them with the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension.

---

## 5. Project Structure

    .
    ├── cmd/                # Application entrypoint
    ├── config/             # Environment configs & DB connection
    ├── entity/             # Domain entities (Post, User, etc.)
    ├── internal/
    │   ├── app/            # Domain-level service, repository, dto, mapper
    │   └── handler/        # Gin handlers (routing)
    ├── pkg/                # Common utils, logger, transaction helpers
    ├── run-server.sh       # Server runner script
    ├── docker-compose.yml  # MySQL docker-compose config
    ├── test.http           # API test file
    └── main.go             # Main entrypoint

---

## 6. Notes
- Database schema is auto-generated via **GORM AutoMigrate**.
- Default charset: `utf8mb4_unicode_ci`.
