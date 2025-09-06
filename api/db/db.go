package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

// GetDB lazily initializes the SQL connection using DB_URL.
// Exits the process with a clear error if DB_URL is missing or the connection fails.
func GetDB() *sql.DB {
	if conn != nil {
		return conn
	}
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL is not set; cannot initialize database connection")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}
	conn = db
	return conn
}
