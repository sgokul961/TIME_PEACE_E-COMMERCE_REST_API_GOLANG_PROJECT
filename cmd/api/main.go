package main

import (
	"log"

	"github.com/joho/godotenv"
	"gokul.go/cmd/api/docs"
	"gokul.go/pkg/config"
	"gokul.go/pkg/di"
)

// @title Go + Gin E-Commerce API
// @version 1.0.0
// @description
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:3000
// @BasePath /
// @query.collection.format multi

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, configErr := config.LoadConfig()

	if configErr != nil {

		log.Fatal("cannot load config:", configErr)

	}

	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "TIME-PEACE"
	docs.SwaggerInfo.Description = "STEP INTO THE WORLD OF FASION WITH UNIQUE STATEMENT PIECS.."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "gokulsajeev.shop"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {
		server.Start()
	}

}
