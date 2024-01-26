// user.go
package routes

import (
	"log"

	"github.com/Zaw-Thet-Paing/API_V1/controllers"
	"github.com/Zaw-Thet-Paing/API_V1/databases"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo) {

	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	userController := controllers.UserController{DB: db}

	e.GET("/users", userController.GetUsers)
	e.GET("/users/:id", userController.GetUser)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
}
