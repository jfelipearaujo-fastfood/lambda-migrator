package loader

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jfelipearaujo-org/lambda-migrator/internal/environment"
	"github.com/stretchr/testify/assert"
)

func cleanEnv() {
	os.Unsetenv("DB_PRODUCTS_URL")
	os.Unsetenv("DB_PRODUCTS_NAME")
	os.Unsetenv("DB_ORDERS_URL")
	os.Unsetenv("DB_ORDERS_NAME")
	os.Unsetenv("DB_PAYMENTS_URL")
	os.Unsetenv("DB_PAYMENTS_NAME")
	os.Unsetenv("DB_PRODUCTIONS_URL")
	os.Unsetenv("DB_PRODUCTIONS_NAME")
	os.Unsetenv("DB_CUSTOMERS_URL")
	os.Unsetenv("DB_CUSTOMERS_NAME")
}

func TestGetEnvironment(t *testing.T) {
	t.Run("Should load environment variables", func(t *testing.T) {
		// Arrange
		t.Setenv("DB_PRODUCTS_URL", "db://host:1234")
		t.Setenv("DB_PRODUCTS_NAME", "product_db")
		t.Setenv("DB_ORDERS_URL", "db://host:1234")
		t.Setenv("DB_ORDERS_NAME", "order_db")
		t.Setenv("DB_PAYMENTS_URL", "db://host:1234")
		t.Setenv("DB_PAYMENTS_NAME", "payment_db")
		t.Setenv("DB_PRODUCTIONS_URL", "db://host:1234")
		t.Setenv("DB_PRODUCTIONS_NAME", "production_db")
		t.Setenv("DB_CUSTOMERS_URL", "db://host:1234")
		t.Setenv("DB_CUSTOMERS_NAME", "customer_db")

		expected := &environment.Config{
			DbProductsConfig: &environment.DatabaseConfig{
				Name: "product_db",
				Url:  "db://host:1234",
			},
			DbOrdersConfig: &environment.DatabaseConfig{
				Name: "order_db",
				Url:  "db://host:1234",
			},
			DbPaymentsConfig: &environment.DatabaseConfig{
				Name: "payment_db",
				Url:  "db://host:1234",
			},
			DbProductionsConfig: &environment.DatabaseConfig{
				Name: "production_db",
				Url:  "db://host:1234",
			},
			DbCustomersConfig: &environment.DatabaseConfig{
				Name: "customer_db",
				Url:  "db://host:1234",
			},
		}

		// Act
		env, err := NewLoader().GetEnvironment(context.Background())

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, env)
		assert.Equal(t, expected, env)
	})

	t.Run("Should return error if a required variable is not set", func(t *testing.T) {
		// Arrange
		t.Setenv("DB_PRODUCTS_URL", "db://host:1234")
		t.Setenv("DB_ORDERS_URL", "db://host:1234")
		t.Setenv("DB_PAYMENTS_URL", "db://host:1234")
		os.Unsetenv("DB_PRODUCTIONS_URL")

		// Act
		env, err := NewLoader().GetEnvironment(context.Background())

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})
}

func TestGetEnvironmentFromFile(t *testing.T) {
	t.Run("Should load environment variables from file", func(t *testing.T) {
		// Arrange
		cleanEnv()

		expected := &environment.Config{
			DbProductsConfig: &environment.DatabaseConfig{
				Name: "product_db",
				Url:  "db://host:1234",
			},
			DbOrdersConfig: &environment.DatabaseConfig{
				Name: "order_db",
				Url:  "db://host:1234",
			},
			DbPaymentsConfig: &environment.DatabaseConfig{
				Name: "payment_db",
				Url:  "db://host:1234",
			},
			DbProductionsConfig: &environment.DatabaseConfig{
				Name: "production_db",
				Url:  "db://host:1234",
			},
			DbCustomersConfig: &environment.DatabaseConfig{
				Name: "customer_db",
				Url:  "db://host:1234",
			},
		}

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "test.env"))

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, env)
		assert.Equal(t, expected, env)
	})

	t.Run("Should return error if a required variable is not set", func(t *testing.T) {
		// Arrange
		cleanEnv()

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "test_err.env"))

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})

	t.Run("Should return error try to load from an invalid file", func(t *testing.T) {
		// Arrange
		cleanEnv()

		// Act
		env, err := NewLoader().GetEnvironmentFromFile(context.Background(), filepath.Join("testdata", "non_exists.env"))

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
	})
}
