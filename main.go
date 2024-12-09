package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"spoti-card.com/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Cannot load env with error")
		fmt.Println(err)
	}

	server := http.NewServeMux()
	router.SpotifyRoute(server)

	servErr := http.ListenAndServe("localhost:3030", server)

	if servErr != nil {
		fmt.Println(servErr)
	}
}
