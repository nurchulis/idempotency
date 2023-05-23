package config

import (
	"os"
	"strconv"
)

type (
	Config struct {
		AppName string
		AppPort int
		RedisConfig
		DBConfig
	}

	RedisConfig struct {
		Addr     string
		Port     int
		Password string
		DB       int
	}

	DBConfig struct {
		Host        string
		Port        int
		User        string
		Password    string
		DBName      string
		Driver      string
		AutoMigrate bool
		AutoSeed    bool
	}
)

func Initialize() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "go-boilerplate"),
		AppPort: getEnvAsInt("APP_PORT", 8080),
		RedisConfig: RedisConfig{
			Addr:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		DBConfig: DBConfig{
			Host:        getEnv("DB_HOST", "localhost"),
			Port:        getEnvAsInt("DB_PORT", 3306),
			User:        getEnv("DB_USER", "mysql"),
			Password:    getEnv("DB_PASSWORD", ""),
			DBName:      getEnv("DB_NAME", "db_idempotency"),
			Driver:      getEnv("DB_DRIVER", "mysql"),
			AutoMigrate: getEnvAsBool("DB_AUTO_MIGRATE", true),
			AutoSeed:    getEnvAsBool("DB_AUTO_SEED", true),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}

	if nextValue := os.Getenv(key); nextValue != "" {
		return nextValue
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}
