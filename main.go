package main

import (
	"github.com/labstack/echo/v4"

	"github.com/vnunicorn/testmaster/internal/database"
	"github.com/vnunicorn/testmaster/internal/user_service"
)

func main() {
	database.InitDb()

	e := echo.New()

	// Routes
	e.POST("/users", user_service.CreateUser)
	e.GET("/users/:id", user_service.GetUser)
	e.PUT("/users/:id", user_service.UpdateUser)
	e.DELETE("/users/:id", user_service.DeleteUser)

	// Start server at localhost:3000
	e.Logger.Fatal(e.Start(":3000"))
}
