package main

import (
	"OrderManagement/config"
	"OrderManagement/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("environment.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.LoadDatabaseConfig()
	server := server.NewServer()
	if err := server.Start(); err != nil {
		panic(err)
	}

}
