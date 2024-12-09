package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"spoti-card.com/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load env with error")
		fmt.Println(err)
	}

	server := http.NewServeMux()
	router.SpotifyRoute(server)

	err = http.ListenAndServe("localhost:3030", server)

	if err != nil {
		log.Fatal("Cannot listen")
		fmt.Println(err)
	}
}
