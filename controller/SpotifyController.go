package controller

import (
	"net/http"

	"spoti-card.com/data"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method != "get" {
		response.Write([]byte("Error"))
	}

	repo := data.InitSpotifyRepository()

	result := repo.GetSpotifyCard()
	response.Write([]byte(result))
}
