package adapters

import (
	"testing"
	"os"
)

var configPost = PostgresConfig{
	Username: "root",
	Password: "root",
	Host:     "localhost",
	Port:     5432,
	Database: "todo",
}

func TestConnectPostgres(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		db, err := ConnectPostgres(configPost)
		if err != nil {
			t.Errorf("ConnectPostgres returned an unexpected error: %v", err)
		}
		defer db.Close()

		if db == nil {
			t.Error("ConnectPostgres returned a nil database")
		}
	})
}

func TestConnectPostgresError(t *testing.T) {
	// Test case where Postgres connection fails
	config := PostgresConfig{
		Username: "invalid",
		Password: "invalid",
		Host: "localhost",
		Port: 5432,
		Database: "test",
	}
	_, err := ConnectPostgres(config)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNewPostgresAdapter(t *testing.T) {
	db, err := ConnectPostgres(configPost)

	if err != nil {
		// handle the error
		t.Errorf("Error getting DB")
	}
	// Call the NewPostgresAdapter function
	adapter := NewPostgresAdapter(db)

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


func TestLoadPostgresConfigFromEnv(t *testing.T) {
	// Test case where an environment variable is not set
	os.Unsetenv("DB_POSTGRES_USERNAME")
	_, err := LoadPostgresConfigFromEnv(nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case where all required environment variables are set
	os.Setenv("DB_POSTGRES_USERNAME", "root")
	os.Setenv("DB_POSTGRES_PASSWORD", "password")
	os.Setenv("DB_POSTGRES_HOST", "localhost")
	os.Setenv("DB_POSTGRES_PORT", "5432")
	os.Setenv("DB_POSTGRES_DATABASE", "test")
	config, err := LoadPostgresConfigFromEnv(nil)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if config.Username != "root" || config.Password != "password" || config.Host != "localhost" || config.Port != 5432 || config.Database != "test" {
		t.Errorf("Expected config to be {Username: root, Password: password, Host: localhost, Port: 5432, Database: test}, got %v", config)
	}
}
