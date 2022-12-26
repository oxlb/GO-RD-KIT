# Initialize the project and install dependencies
.PHONY: setup
setup:
	go mod init github.com/oxlb/GO-RD-KIT
	go get github.com/jinzhu/gorm
	go get github.com/jinzhu/gorm/dialects/mysql
	go get github.com/joho/godotenv

.PHONY: test
test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

# Format code
.PHONY: format
format:
	go fmt ./...