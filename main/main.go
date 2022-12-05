package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"main/model"
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

	resp, err := http.Get("https://imdb-api.com/en/API/Top250TVs/" + api_key)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var topMovies model.TopMovies
	json.Unmarshal(bodyBytes, &topMovies)

	//Add DATA to Redis

}
