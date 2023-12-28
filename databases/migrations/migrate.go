package migrations

import (
	"fmt"
	"log"

	"github.com/Zaw-Thet-Paing/API_V1/databases"
	"github.com/Zaw-Thet-Paing/API_V1/models/entities"
)

func Migrate(action string, model_name string) {
	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	if action == "migrate" {
		if model_name == "all" {
			db.AutoMigrate(&entities.User{})
			fmt.Println("All models migrate success...")
		} else if model_name == "user" {
			db.AutoMigrate(&entities.User{})
			fmt.Println("User model migrate success...")
		} else {
			fmt.Printf("No match model name with %s", model_name)
		}
	} else if action == "drop" {
		if model_name == "all" {
			db.Migrator().DropTable(&entities.User{})
			fmt.Println("All models drop success...")
		} else if model_name == "user" {
			db.Migrator().DropTable(&entities.User{})
			fmt.Println("User model drop success...")
		} else {
			fmt.Printf("No match model name with %s", model_name)
		}
	} else {
		fmt.Println("Unknown action")
	}

}
