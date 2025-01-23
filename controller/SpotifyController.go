package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"spoti-card.com/controller/utils"
	"spoti-card.com/data"
	"spoti-card.com/presentation"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	tokenRepo := data.TokenRepository()

	accessToken, err := tokenRepo.FetchAccessToken()
	if err != nil {
		log.Default().Printf("access token" + err.Error())
		utils.HandleError(err, response)
		return
	}

	clientToken, err := tokenRepo.FetchClientToken(accessToken.ClientId)
	if err != nil {
		log.Default().Printf("client token" + err.Error())
		utils.HandleError(err, response)
		return
	}

	fmt.Printf("New access token %s\n", accessToken.AccessToken)

	trackRepo := data.TrackRepository(*accessToken, *clientToken)

	playbackState, err := trackRepo.GetPlaybackState()
	if err != nil {
		log.Default().Printf("playback state" + err.Error())
		utils.HandleError(err, response)
		return
	}

	trackId := strings.ReplaceAll(playbackState.PlayerState.Track.Uri, "spotify:track:", "")

	trackResult, err := trackRepo.GetTrackById(trackId)
	if err != nil {
		log.Default().Printf("track" + err.Error())
		utils.HandleError(err, response)
		return
	}

	cardModel := presentation.SpoticardModel{
		Track:     *trackResult,
		IsPlaying: playbackState.PlayerState.IsPlaying && !playbackState.PlayerState.IsPaused,
	}
	card := presentation.SpotifyCard(cardModel)

	fmt.Println("All request succeeded")

	response.Header().Add("Content-Type", "text/html")
	response.Write([]byte(card))
}
