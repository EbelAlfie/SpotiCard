package controller

import (
	"net/http"

	"spoti-card.com/data"
	"spoti-card.com/presentation"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	tokenRepo := data.TokenRepository()

	accessToken, err := tokenRepo.FetchAccessToken()
	if err != nil {
		return
	}

	clientToken, err := tokenRepo.FetchClientToken(accessToken.ClientId)
	if err != nil {
		return
	}

	trackRepo := data.TrackRepository(*accessToken, *clientToken)

	playbackState, err := trackRepo.GetDeviceState()
	if err != nil {
		return
	}

	trackResult, err := trackRepo.GetTrackById(playbackState.Track.Uri)
	if err != nil {
		response.Write([]byte(err.Error()))
		return
	}

	card := presentation.SpotifyCard()

	response.Write([]byte(card))
}
