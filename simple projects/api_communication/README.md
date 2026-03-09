# Stock API - Docker Setup

This is a Go-based REST API for managing stock inventory using the Gin framework.

## Prerequisites

- Docker installed on your system
- Docker Compose (optional, but recommended)

## Project Structure

```
api_communication/
├── Dockerfile           # Multi-stage Docker build configuration
├── .dockerignore       # Files to exclude from Docker build
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksums
├── main.go             # Main application entry point
├── stock.go            # Stock/inventory handlers
└── hello.go            # Hello/hi handlers
```

## API Endpoints

- `GET /hi` - Calls the /hello endpoint internally
- `GET /hello` - Returns a greeting message
- `POST /order` - Process an order
- `POST /update_inventory` - Update inventory after order
- `GET /inventory` - Get current inventory status

## Docker Build & Run Instructions

### Option 1: Using Docker Compose (Recommended)

**Build and start the container:**
```bash
cd api_communication
docker-compose up --build
```

**Run in detached mode:**
```bash
docker-compose up -d
```

**Stop the container:**
```bash
docker-compose down
```

**View logs:**
```bash
docker-compose logs -f
```

### Option 2: Using Docker Commands

**Build the Docker image:**
```bash
cd api_communication
docker build -t stock-api:latest .
```

**Run the container:**
```bash
docker run -d -p 8080:8080 --name stock-api stock-api:latest
```

**Stop the container:**
```bash
docker stop stock-api
```

**Remove the container:**
```bash
docker rm stock-api
```

**View logs:**
```bash
docker logs -f stock-api
```

## Testing the API

Once the container is running, test the endpoints:

**Check inventory:**
```bash
curl http://localhost:8080/inventory
```

**Test hello endpoint:**
```bash
curl http://localhost:8080/hello
```

**Test hi endpoint (calls hello internally):**
```bash
curl http://localhost:8080/hi
```

**Place an order:**
```bash
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{"product":"mouse","quantity":5}'
```

**Update inventory:**
```bash
curl -X POST http://localhost:8080/update_inventory \
  -H "Content-Type: application/json" \
  -d '{"product":"mouse","quantity":5}'
```

## Docker Image Details

- **Base Image:** Alpine Linux (minimal footprint)
- **Go Version:** 1.25
- **Build Type:** Multi-stage build for optimized image size
- **Port:** 8080
- **Architecture:** linux/amd64

## Troubleshooting

**Container won't start:**
```bash
docker logs stock-api
```

**Port already in use:**
```bash
# Change the port mapping in docker-compose.yml or use:
docker run -d -p 8081:8080 --name stock-api stock-api:latest
```

**Rebuild after code changes:**
```bash
docker-compose up --build
# or
docker build --no-cache -t stock-api:latest .
```

## Notes

- The application uses in-memory storage, so inventory resets on container restart
- The `/update_inventory` endpoint makes an internal call to `/order` on localhost:8080
- For production use, consider adding environment variables for configuration
- Consider using a persistent database instead of in-memory storage