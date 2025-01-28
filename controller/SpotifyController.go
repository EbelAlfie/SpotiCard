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
	response.Header().Set("Content-Type", "image/svg+xml")

	requestParam := request.URL.Query()

	code := requestParam.Get("code")
	if code == "" {
		utils.HandleError(response, errors.New("who are you?"), http.StatusBadRequest)
		return 
	}

	tokenRepository := data.TokenRepository(code)

	tokenData, err := tokenRepository.FetchAccessToken()
	if err != nil {
		fmt.Printf("token %s", err)
		http.Redirect(response, request, "http://localhost:3031/login", http.StatusTemporaryRedirect)
		//utils.HandleError(response, err, http.StatusBadGateway)
		return 
	}

	fmt.Println(tokenData.AccessToken)

	trackRepository := data.TrackRepository(tokenData.AccessToken)

	trackData, err := trackRepository.GetRecentlyPlayed()
	if httpErr, isType := err.(*entity.HttpError); isType {
		trackData, err := trackRepository.GetRecentlyPlayed()
	}

	if err != nil {
		fmt.Printf("playback %s", err)
		utils.HandleError(response, err, http.StatusBadGateway)
		return
	}

	fmt.Println("Track")
	fmt.Println(trackData)
	// fmt.Println(trackData.IsPlaying)

	cardModel := presentation.SpoticardModel {
		Track: *trackData,
		IsPlaying: false, //trackData.IsPlaying,
	}
	spotiCard := presentation.SpotifyCard(cardModel)
	
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(spotiCard))
}
