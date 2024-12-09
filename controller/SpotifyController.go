package controller

import (
	"log"
	"net/http"
	"strings"

	"spoti-card.com/data"
	"spoti-card.com/presentation"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	tokenRepo := data.TokenRepository()

	accessToken, err := tokenRepo.FetchAccessToken()
	if err != nil {
		log.Default().Printf("access token" + err.Error())
		response.Write([]byte(err.Error()))
		return
	}

	clientToken, err := tokenRepo.FetchClientToken(accessToken.ClientId)
	if err != nil {
		log.Default().Printf("client token" + err.Error())
		response.Write([]byte(err.Error()))
		return
	}

	trackRepo := data.TrackRepository(*accessToken, *clientToken)

	playbackState, err := trackRepo.GetPlaybackState()
	if err != nil {
		log.Default().Printf("playback state" + err.Error())
		response.Write([]byte(err.Error()))
		return
	}

	trackId := strings.ReplaceAll(playbackState.PlayerState.Track.Uri, "spotify:track:", "")

	trackResult, err := trackRepo.GetTrackById(trackId)
	if err != nil {
		log.Default().Printf("track" + err.Error())
		response.Write([]byte(err.Error()))
		return
	}

	card := presentation.SpotifyCard(*trackResult)

	response.Header().Add("Content-Type", "text/html")
	response.Write([]byte(card))
}
