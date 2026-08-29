package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey              string
	ExpirationHours        time.Duration
	RefreshExpirationHours time.Duration
}

type AWSConfig struct {
	Region     string
	AccessKey  string
	SecretKey  string
	BucketName string
	Endpoint   string
}

type UploadConfig struct {
	Path    string
	MaxSize int64
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	jwtExpirationHours, _ := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	jwtRefreshExpirationHours, _ := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRATION_HOURS", "168"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("UPLOAD_MAX_SIZE", "10485760"), 10, 64)

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "ecommerce"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:              getEnv("JWT_SECRET_KEY", "your-secret-key"),
			ExpirationHours:        time.Duration(jwtExpirationHours),
			RefreshExpirationHours: time.Duration(jwtRefreshExpirationHours),
		},
		AWS: AWSConfig{
			Region:     getEnv("AWS_REGION", "us-east-1"),
			AccessKey:  getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretKey:  getEnv("AWS_SECRET_ACCESS_KEY", ""),
			BucketName: getEnv("AWS_BUCKET_NAME", ""),
			Endpoint:   getEnv("AWS_ENDPOINT", ""),
		},
		Upload: UploadConfig{
			Path:    getEnv("UPLOAD_PATH", "./uploads"),
			MaxSize: maxUploadSize,
		},
	}, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
