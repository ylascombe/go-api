package database

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/ylascombe/go-api/models"
)

func ListManifests(db *sql.DB) {

	var (
		id       int
		content string
	)

	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT id FROM manifest")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var m models.Manifest
	for rows.Next() {
		err := rows.Scan(&id, &content)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(m, "id =", id, ", ", "content =", content)
	}
}