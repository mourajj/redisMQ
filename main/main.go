package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("variables.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_key := os.Getenv("API_KEY")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

}
