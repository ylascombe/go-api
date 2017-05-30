package database

import (
	"fmt"
	"os"
	"log"
	"database/sql"
)

func SetupDB(db *sql.DB) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s port=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPORT"),
	)
	log.Println("dsn =", dsn)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}