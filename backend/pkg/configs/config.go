package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db  DbConfig
	Ssl SSLConfig
	Jwt JWTConfig
}

type DbConfig struct {
	Dsn string
}

type SSLConfig struct {
	SSLCertPath string
	SSLKeyPath  string
}

type JWTConfig struct {
	SecretKey  string
	AccessTTL  string
	RefreshTTL string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Ssl: SSLConfig{
			SSLCertPath: os.Getenv("SSL_CERT_PATH"),
			SSLKeyPath:  os.Getenv("SSL_KEY_PATH"),
		},
		Jwt: JWTConfig{
			SecretKey:  os.Getenv("SECRET_KEY"),
			AccessTTL:  os.Getenv("JWT_TTL"),
			RefreshTTL: os.Getenv("JWT_REFRESH_TTL"),
		},
	}
}
