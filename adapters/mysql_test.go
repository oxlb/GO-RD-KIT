package adapters

import (
	"testing"
	"os"
)

var config = MySQLConfig{
	Username: "root",
	Password: "root",
	Host:     "localhost",
	Port:     3306,
	Database: "todo",
}

func TestConnectMySQL(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, err := ConnectMySQL(config)
		if err != nil {
			t.Errorf("ConnectMySQL returned an unexpected error: %v", err)
		}
		defer db.Close()

		if db == nil {
			t.Error("ConnectMySQL returned a nil database")
		}
	})
}

func TestConnectMySQLError(t *testing.T) {
	// Test case where MySQL connection fails
	config := MySQLConfig{
		Username: "invalid",
		Password: "invalid",
		Host: "localhost",
		Port: 3306,
		Database: "test",
	}
	_, err := ConnectMySQL(config)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNewAdapter(t *testing.T) {
	db, err := ConnectMySQL(config)

	if err != nil {
		// handle the error
		t.Errorf("Error getting DB")
	}
	// Call the NewAdapter function
	adapter := NewMySQLAdapter(db)

	// Assert that the returned adapter is not nil
	if adapter == nil {
		t.Errorf("Expected adapter to be non-nil, got nil")
	}

	// Assert that the adapter's DB field is set to the database connection
	if adapter.DB != db {
		t.Errorf("Expected adapter.DB to be set to the database connection, got %v", adapter.DB)
	}

	// Close the database object.
	defer db.Close()
}


func TestLoadMySQLConfigFromEnv(t *testing.T) {
	// Test case where an environment variable is not set
	os.Unsetenv("DB_MYSQL_USERNAME")
	_, err := LoadMySQLConfigFromEnv(nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case where all required environment variables are set
	os.Setenv("DB_MYSQL_USERNAME", "root")
	os.Setenv("DB_MYSQL_PASSWORD", "password")
	os.Setenv("DB_MYSQL_HOST", "localhost")
	os.Setenv("DB_MYSQL_PORT", "3306")
	os.Setenv("DB_MYSQL_DATABASE", "test")
	config, err := LoadMySQLConfigFromEnv(nil)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if config.Username != "root" || config.Password != "password" || config.Host != "localhost" || config.Port != 3306 || config.Database != "test" {
		t.Errorf("Expected config to be {Username: root, Password: password, Host: localhost, Port: 3306, Database: test}, got %v", config)
	}
}
