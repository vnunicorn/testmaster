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

type User struct {
	Username    string
	Mobile      string
	Email       string
	Password    string
	First_name  string
	Middle_name string
	Last_name   string
}

func AddUser(user *User) (int64, error) {
	sql := fmt.Sprintf("INSERT INTO testmaster.users(username, mobile, email, password, first_name, middle_name, last_name) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')",
		user.Username, user.Mobile, user.Email, user.Password, user.First_name, user.Middle_name, user.Last_name)
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)

	return lastId, err
}
