package main

import (
	"LogiGO/cmd/database"
	"LogiGO/cmd/models"
	"LogiGO/cmd/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	models.InitValidators()
	database.ConnectDB()
	routes.HandlerRequest()
}
