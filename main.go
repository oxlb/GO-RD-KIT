package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/oxlb/GO-RD-KIT/adapters"
	"github.com/oxlb/GO-RD-KIT/models"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load MySQL configuration from environment variables or pass conf in the param
	mysqlConfig, err := adapters.LoadMySQLConfigFromEnv(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Connect to MySQL database
	mysqlDB, err := adapters.ConnectMySQL(mysqlConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	var todos []models.Todos
	mysqlDB.Find(&todos)

	for _, todo := range todos {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", todo.ID, todo.Title, todo.Description, todo.Completed)
	}

	if len(todos) == 0 {
		fmt.Println("No rows found")
	}

}
