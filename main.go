package main

import (
	"fmt"
	"net/http"

	"spoti-card.com/router"
)

func main() {
	server := http.NewServeMux()
	router.SpotifyRoute(server)

	servErr := http.ListenAndServe("localhost:3030", server)

	if servErr != nil {
		fmt.Println(servErr)
	}
}
