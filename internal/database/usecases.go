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

//thêm mới user
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

// lấy ra tất cả user
func GetAllUsers() ([]User, error) {
	var sliceUsers []User
	result, err := db.Query("SELECT * FROM testmaster.users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var u User
		_ = result.Scan(&u.Username, &u.Mobile, &u.Email, &u.Password, &u.First_name, &u.Middle_name, &u.Last_name)
		sliceUsers = append(sliceUsers, u)
	}
	return sliceUsers, err
}

// lấy User theo id
func GetUserByID(id string) (User, error) {

	getDB, err := db.Query("SELECT * FROM testmaster.users WHERE id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	var u User
	fmt.Println(getDB.Columns())
	_ = getDB.Scan(&u.Username, &u.Mobile, &u.Email, &u.Password, &u.First_name, &u.Middle_name, &u.Last_name)

	fmt.Printf("Select username by id: %s\n", id)

	return u, err

}

//Update user by id
func UpdateUserByID(user *User, id string) error {

	res, err := db.Prepare("UPDATE testmaster.users SET username=?, mobile=?, email=?, password id=?, first_name=?,middle_name=?, last_name=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	res.Exec(user.Username, user.Mobile, user.Email, user.Password, user.First_name, user.Middle_name, user.Last_name, id)
	fmt.Printf("Update username by id successfully")
	return err

}

//Delete user by id
func DeleteUserById(id string) error {
	res, err := db.Prepare("DELETE FROM testmaster.users WHERE id=?")
	if err != nil {
		panic(err)
	}
	res.Exec(id)
	fmt.Printf("Delete username by id successfully")
	return nil
}
