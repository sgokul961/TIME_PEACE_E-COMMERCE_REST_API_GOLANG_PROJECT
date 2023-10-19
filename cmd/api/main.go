package main

import (
	"log"

	"github.com/joho/godotenv"
	"gokul.go/pkg/config"
	"gokul.go/pkg/di"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, configErr := config.LoadConfig()

	if configErr != nil {
		log.Fatal("cannot load config:", configErr)
	}
	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {
		server.Start()
	}

}
