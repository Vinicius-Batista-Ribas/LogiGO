package main

import (
	"LogiGO/cmd/database"
	"LogiGO/cmd/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	database.ConnectDB()
	routes.HandlerRequest()
}
