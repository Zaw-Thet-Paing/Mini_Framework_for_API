package routes

import (
	"github.com/Zaw-Thet-Paing/API_V1/controllers"
	// "net/http"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	//user controller
	user := controllers.UserController{}
	e.GET("/users", user.GetUsers)
	e.GET("/users/:id", user.GetUser)
	e.POST("/users", user.CreateUser)
	e.PUT("/users/:id", user.UpdateUser)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "<h1>Hello Go</h1>")
	// })
}
