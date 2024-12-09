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

	playbackState, err := trackRepo.GetPlaybackState()
	if err != nil {
		return
	}

	trackResult, err := trackRepo.GetTrackById(playbackState.Track.Uri)
	if err != nil {
		return
	}

	card := presentation.SpotifyCard(*trackResult)

	response.Write([]byte(card))
}
