package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	// replace with your actual module path
	"backend/internal/db"
	"backend/internal/router"
)

func main() {
	// Read DATABASE_URL from env
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	// Connect to Postgres
	config, err := pgx.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("failed to parse DB config: %v", err)
	}
	
	// Use stdlib to get sql.DB
	sqlDB := stdlib.OpenDB(*config)
	defer sqlDB.Close()

	// Initialize sqlc queries
	queries := db.New(sqlDB)

	// Build router with all routes & middleware
	r := router.NewRouter(queries)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
