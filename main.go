package main

import (
	"fmt"
	"os"

	"github.com/Zaw-Thet-Paing/API_V1/databases/migrations"
	"github.com/Zaw-Thet-Paing/API_V1/routes"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {

	args_len := len(os.Args[1:])

	if args_len == 0 {
		e := echo.New()
		routes.Setup(e)
		e.Logger.Fatal(e.Start(":1234"))
	} else if args_len == 2 {
		if os.Args[1] == "migrate" || os.Args[1] == "drop" {
			migrations.Migrate(os.Args[1], os.Args[2])
		} else {
			fmt.Println("Unknown arguments")
		}
	} else {
		fmt.Println("Error: Arguments")
	}
}
