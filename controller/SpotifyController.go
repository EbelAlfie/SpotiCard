package controller

import (
	"errors"
	"fmt"
	"net/http"

	"spoti-card.com/controller/utils"
	"spoti-card.com/data"
	"spoti-card.com/domain/entity"
	"spoti-card.com/presentation"
)

func SpotifyController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "image/svg+xml")
	response.Header().Set("Cache-Control", "max-age=1 ,s-maxage=1")

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

	var trackData *entity.TrackEntity
	var isPlaying bool
	playbackState, err := trackRepository.GetPlaybackState()

	if playbackState != nil {
		trackData = &playbackState.Track
		isPlaying = playbackState.IsPlaying
	}

	if httpErr, isType := err.(*entity.HttpError); isType && httpErr.StatusCode == 204 {
		trackData, err = trackRepository.GetRecentlyPlayed()
		isPlaying = false
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
		IsPlaying: isPlaying, //trackData.IsPlaying,
	}
	spotiCard := presentation.SpotifyCard(cardModel)
	
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(spotiCard))
}
