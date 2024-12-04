package controller

import (
	"net/http"

	"spoti-card.com/data"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	repo := data.InitSpotifyRepository()

	result := repo.GetSpotifyCard()

	response.Write([]byte(result))
}
