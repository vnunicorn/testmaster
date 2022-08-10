package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDb() {
	cfg := mysql.Config{
		User:   "root",       // Set Username Here
		Passwd: "testmaster", // Set Password Here
		Net:    "tcp",
		Addr:   "127.0.0.1:3308",
		DBName: "testmaster", // Set Database Name Here
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Println("can not connect")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("can not ping")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
