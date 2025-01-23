package controller

import (
	"errors"
	"fmt"
	"net/http"

	"spoti-card.com/controller/utils"
	"spoti-card.com/data"
	"spoti-card.com/presentation"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	requestParam := request.URL.Query()

	code := requestParam.Get("code")
	if code == "" {
		utils.HandleError(response, errors.New("who are you?"), http.StatusBadRequest)
		return 
	}

	tokenRepository := data.TokenRepository(code)

	tokenData, err := tokenRepository.FetchAccessToken()
	if err != nil {
		utils.HandleError(response, err, http.StatusBadGateway)
		return 
	}

	trackRepository := data.TrackRepository(tokenData.AccessToken)

	playbackState, err := trackRepository.GetPlaybackState()
	if err != nil {
		fmt.Print(err)
		utils.HandleError(response, err, http.StatusBadGateway)
		return 
	}

	cardModel := presentation.SpoticardModel {
		Track: playbackState.Track,
		IsPlaying: playbackState.IsPlaying,
	}
	spotiCard := presentation.SpotifyCard(cardModel)
	
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "image/svg+xml")
	response.Write([]byte(spotiCard))
}
