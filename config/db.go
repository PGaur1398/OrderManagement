package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
}

var dbConfig DatabaseConfig

func LoadDatabaseConfig() {
	dbConfig = DatabaseConfig{
		Port:     os.Getenv("APP_PORT"),
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("DATABASE_NAME"),
	}
}

func GetDBConfig() DatabaseConfig {
	return dbConfig
}

func (dbconf DatabaseConfig) GetDBUrl() string {

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbconf.User, dbconf.Password, dbconf.Host, dbConfig.DBName)
}
