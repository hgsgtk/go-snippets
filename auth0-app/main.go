package main

import (
	"github.com/joho/godotenv"
	"./app"
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