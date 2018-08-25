package main

import (
	"./app"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	app.Init()
	StartServer()
}
