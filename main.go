package main

import (
	"github.com/caio-rds/golang-api/src/database"
	"github.com/caio-rds/golang-api/src/routes"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db := database.NewDatabase()
	if err := routes.InitRoutes(db); err != nil {
		log.Fatal(err)
	}
}
