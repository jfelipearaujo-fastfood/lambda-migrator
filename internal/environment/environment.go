package environment

import (
	"context"
)

type DatabaseConfig struct {
	Name string `env:"NAME, required"`
	Url  string `env:"URL, required"`
}

type Config struct {
	DbProductsConfig    *DatabaseConfig `env:",prefix=DB_PRODUCTS_"`
	DbOrdersConfig      *DatabaseConfig `env:",prefix=DB_ORDERS_"`
	DbPaymentsConfig    *DatabaseConfig `env:",prefix=DB_PAYMENTS_"`
	DbProductionsConfig *DatabaseConfig `env:",prefix=DB_PRODUCTIONS_"`
}

type Environment interface {
	GetEnvironmentFromFile(ctx context.Context, fileName string) (*Config, error)
	GetEnvironment(ctx context.Context) (*Config, error)
}
