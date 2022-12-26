package adapters

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	// Import postgres drivers
	_ "github.com/jinzhu/gorm/dialects/postgres" // import postgres 
)

// PostgresConfig represents Postgres configuration.
type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

// LoadPostgresConfigFromEnv loads Postgres configuration from environment variables, or from the provided LoadPostgresConfigFromEnv struct if it is not nil.
func LoadPostgresConfigFromEnv(config *PostgresConfig) (PostgresConfig, error) {
	if config == nil {
		port, err := strconv.Atoi(os.Getenv("DB_POSTGRES_PORT"))
		if err != nil {
			return PostgresConfig{}, err
		}

		config = &PostgresConfig{
			Username: os.Getenv("DB_POSTGRES_USERNAME"),
			Password: os.Getenv("DB_POSTGRES_PASSWORD"),
			Host:     os.Getenv("DB_POSTGRES_HOST"),
			Port:     port,
			Database: os.Getenv("DB_POSTGRES_DATABASE"),
		}
	}
	
	return *config, nil
}

func ConnectPostgres(config PostgresConfig) (*gorm.DB, error) {
	// Create the connection string
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.Database)
		// Connect to the Postgres database
		db, err := gorm.Open("postgres", connStr)
		if err != nil {
			return nil, fmt.Errorf("Error connecting to Postgres: %v", err)
		}
	
		return db, nil
}

func NewPostgresAdapter(db *gorm.DB) *PostgresAdapter {
	return &PostgresAdapter{
		DB: db,
	}
}

type PostgresAdapter struct {
	DB *gorm.DB
}
