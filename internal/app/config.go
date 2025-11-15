package app

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort        int           `env:"HTTP_PORT" envDefault:"8080"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`

	DBDriver   string `env:"DB_DRIVER" envDefault:"postgres"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     int    `env:"DB_PORT,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`

	MigrationsPath string        `env:"MIGRATIONS_PATH" envDefault:"./migrations"`
	HTTPTimeout    time.Duration `env:"HTTP_TIMEOUT" envDefault:"5s"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("%s: %w", msgEnvLoadFail, err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", msgEnvParseFail, err)
	}

	return cfg, nil
}

func (c *Config) DSN() (string, error) {
	switch c.DBDriver {
	case "postgres":
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode,
		), nil
	default:
		return "", fmt.Errorf("%s: %s", msgUnsupportedDriver, c.DBDriver)
	}
}
