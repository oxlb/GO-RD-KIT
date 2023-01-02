[![Test](https://github.com/oxlb/GO-RD-KIT/actions/workflows/main.yml/badge.svg)](https://github.com/oxlb/GO-RD-KIT/actions/workflows/main.yml)

![Image Alt Text](https://raw.githubusercontent.com/oxlb/GO-RD-KIT/main/Go-Repo-Go-RD-KIT%20(1).png)


# GO-RD-KIT
A Go SDK for connecting to and interacting with relational databases, such as MySQL and Postgres.

## Prerequisites

- Go

- Docker Desktop

## Installation

To install the dependencies for this project, run the following command:


``` prompt
make setup
```

This will initialize the Go module and install the required dependencies, including gorm and godotenv.

## Running Tests

To run the tests for this project, use the following command:
``` prompt
make test
```

This will run all the tests in the project and generate a coverage report.

## Formatting Code

To format the code in this project, use the following command:

``` prompt
make format
```

This will run the go fmt command to format the code according to the Go code formatting standards.

## Example Usage

Make sure `Docker Desktop` is running

Use following command

``` prompt
docker-compose up
```

To run the example code, use the following command:

``` prompt
go run main.go
```

This will connect to the MySQL database using the configuration in the .env file, retrieve all rows from the todos table, and print the data to the console.

### MySql

``` GO

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

```

### Postgres

``` GO

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

	posgresConfig, err := adapters.LoadPostgresConfigFromEnv(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Connect to Postgres database
	postgresDB, err := adapters.ConnectPostgres(posgresConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	var todosPostgres []models.Todos
	postgresDB.Find(&todosPostgres)

	for _, todo := range todosPostgres {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", todo.ID, todo.Title, todo.Description, todo.Completed)
	}

	if len(todosPostgres) == 0 {
		fmt.Println("No rows found")
	}
  
}

```
