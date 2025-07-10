package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	dsn := os.Getenv("POSTGRES_DSN")
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Failed to open DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Failed to ping DB:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL")
}
