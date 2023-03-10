package main

import (
	"log"
	"os"
	"sinarmas/kredit-sinarmas/api"
	"sinarmas/kredit-sinarmas/database"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.SetOutput(file)

	db, err := database.SetupDb()
	if err != nil {
		log.Panicln(err.Error())
	}

	server := api.MakeServer(db)
	server.RunServer()
}
