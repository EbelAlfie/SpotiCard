package controller

import (
	"net/http"

	"spoti-card.com/data"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	repo := data.InitSpotifyRepository()

	trackResult, err := repo.GetTrackById("")
	if err != nil {
		response.Write([]byte(err.Error()))
		return
	}

	card := ""

	response.Write([]byte(card))
}
