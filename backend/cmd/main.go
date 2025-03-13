package main

import (
	"backend/api"
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

func createTables() {
	connStr := "user=postgres password=password123 dbname=fullstack sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	sqlScript, err := ioutil.ReadFile("sql/createTables.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlScript))

	db.Close()
}

func main() {
	createTables()
	log.Println("tables created")
	// admin.Launch()
	api.Start()
}
