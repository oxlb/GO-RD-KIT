package adapters

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	// Import MySQL drivers
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLConfig represents MySQL configuration.
type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

// LoadMySQLConfigFromEnv loads MySQL configuration from environment variables, or from the provided MySQLConfig struct if it is not nil.
func LoadMySQLConfigFromEnv(config *MySQLConfig) (MySQLConfig, error) {
	if config == nil {
		port, err := strconv.Atoi(os.Getenv("DB_MYSQL_PORT"))
		if err != nil {
			return MySQLConfig{}, err
		}

		config = &MySQLConfig{
			Username: os.Getenv("DB_MYSQL_USERNAME"),
			Password: os.Getenv("DB_MYSQL_PASSWORD"),
			Host:     os.Getenv("DB_MYSQL_HOST"),
			Port:     port,
			Database: os.Getenv("DB_MYSQL_DATABASE"),
		}
	}

	return *config, nil
}

// ConnectMySQL establishes a connection to the MySQL database.
func ConnectMySQL(config MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// MySQLAdapter represents a MySQL adapter.
type MySQLAdapter struct {
	DB *gorm.DB
}

func NewMySQLAdapter(db *gorm.DB) *MySQLAdapter {
	return &MySQLAdapter{DB: db}
}
