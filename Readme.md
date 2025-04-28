# Go React TypeScript Vite Starter Template

A modern, production-ready starter template for building full-stack web applications with Go, React, TypeScript, and Vite.

## Features

### Backend (Go)
- Clean architecture with separation of concerns
- SQLC for type-safe database queries
- Chi router for HTTP routing
- PostgreSQL database integration
- Docker support
- Environment configuration
- Structured logging

### Frontend (React + TypeScript + Vite)
- Modern React with TypeScript
- Vite for fast development and building
- TailwindCSS for styling
- ESLint for code quality
- Docker support
- API proxy configuration
- Type-safe API calls

## Getting Started

### Prerequisites
- Go 1.23.5 or later
- Node.js 18 or later
- PostgreSQL 15 or later
- Docker (optional)

### Backend Setup
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Copy the environment file:
   ```bash
   cp .env.example .env
   ```

3. Update the `.env` file with your database credentials

4. Install dependencies:
   ```bash
   go mod download
   ```

5. Generate SQLC code:
   ```bash
   sqlc generate
   ```

6. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

## Project Structure

```
.
├── backend/
│   ├── cmd/            # Application entrypoints
│   ├── internal/       # Private application code
│   │   ├── db/        # Database models and queries
│   │   └── router/    # HTTP routing
│   ├── schema.sql     # Database schema
│   └── query.sql      # SQLC queries
└── frontend/
    ├── src/           # Source code
    ├── public/        # Static assets
    └── index.html     # Entry HTML file
```

## Development

- Backend API runs on `http://localhost:8080`
- Frontend dev server runs on `http://localhost:5173`
- API requests from frontend are automatically proxied to the backend

## Building for Production

### Backend
```bash
cd backend
go build -o server cmd/server/main.go
```

### Frontend
```bash
cd frontend
npm run build
```

## Docker Support

Both frontend and backend include Dockerfile configurations for containerized deployment.

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
