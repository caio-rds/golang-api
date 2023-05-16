package main

import (
	"github.com/caio-rds/golang-api/src/controller/routes"
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
	if err := routes.InitRoutes(); err != nil {
		log.Fatal(err)
	}
}
