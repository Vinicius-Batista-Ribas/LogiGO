package main

import (
	"LogiGO/cmd/database"
	"LogiGO/cmd/routes"
)

func main() {
	database.ConnectDB()
	routes.HandlerRequest()
}
