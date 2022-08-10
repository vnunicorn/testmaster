package user_service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/vnunicorn/testmaster/internal/database"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//---------

func CreateUser(c echo.Context) error {
	u := &database.User{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	id, err := database.AddUser(u)
	if err != nil {
		glog.Error("err : %v", err)
	}
	log.Printf("created user, id = %d", id)
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
