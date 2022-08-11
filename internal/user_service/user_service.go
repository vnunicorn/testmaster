package user_service

import (
	"log"
	"net/http"

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
	id := c.Param("id")

	user, err := database.GetUserByID(id)

	if err != nil {
		glog.Error("err : %v", err)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	u := new(database.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	err := database.UpdateUserByID(u, c.Param("id"))

	if err != nil {
		glog.Error("err : %v", err)
	}

	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := database.DeleteUserById(id)
	if err != nil {
		panic(err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func GetUsers(c echo.Context) error {
	res, err := database.GetAllUsers()
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
