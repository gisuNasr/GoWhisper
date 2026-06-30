package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	SSLMode  string `toml:"sslmode"`
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
}

func Load(path string) (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if _, err := os.Stat(path); err == nil && path != ".env" {
		if _, err := toml.DecodeFile(path, &cfg); err != nil {
			return nil, fmt.Errorf("failed to load config from %s: %w", path, err)
		}
	}

	if envPort := os.Getenv("APP_PORT"); envPort != "" {
		if p, err := strconv.Atoi(envPort); err == nil {
			cfg.Server.Port = p
		}
	}
	if envHost := os.Getenv("DB_HOST"); envHost != "" {
		cfg.Database.Host = envHost
	}
	if envDbPort := os.Getenv("DB_PORT"); envDbPort != "" {
		if p, err := strconv.Atoi(envDbPort); err == nil {
			cfg.Database.Port = p
		}
	}
	if envUser := os.Getenv("DB_USER"); envUser != "" {
		cfg.Database.User = envUser
	}
	if envPassword := os.Getenv("DB_PASSWORD"); envPassword != "" {
		cfg.Database.Password = envPassword
	}
	if envName := os.Getenv("DB_NAME"); envName != "" {
		cfg.Database.Name = envName
	}
	if envSSLMode := os.Getenv("DB_SSLMODE"); envSSLMode != "" {
		cfg.Database.SSLMode = envSSLMode
	}

	return &cfg, nil
}
