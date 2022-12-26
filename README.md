# GO-RD-KIT
A Go SDK for connecting to and interacting with relational databases, such as MySQL and Postgres.

## Prerequisites

- Go

- Docker Desktop

`Installation`

To install the dependencies for this project, run the following command:


`make setup`

This will initialize the Go module and install the required dependencies, including gorm and godotenv.

## Running Tests

To run the tests for this project, use the following command:

`make test`

This will run all the tests in the project and generate a coverage report.

## Formatting Code

To format the code in this project, use the following command:

`make format`

This will run the go fmt command to format the code according to the Go code formatting standards.

## Example Usage

Make sure `Docker Desktop` is running

run following command

`docker-compose up`

To run the example code, use the following command:

`go run main.go`

This will connect to the MySQL database using the configuration in the .env file, retrieve all rows from the todos table, and print the data to the console.
