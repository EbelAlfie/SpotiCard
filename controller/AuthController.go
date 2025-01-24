package controller

import (
	"net/http"
	"net/url"
	"os"
)

func AuthController(response http.ResponseWriter, request *http.Request) {
	rawUrl := "https://accounts.spotify.com/authorize?"
	
	authUrl, _ := url.Parse(rawUrl)

	clientID := os.Getenv("CLIENT_ID")
	redirectTarg := "http://localhost:3031"
	scope := "user-read-currently-playing user-read-playback-state user-read-recently-played"

	params := authUrl.Query()
	params.Add("client_id", clientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirectTarg)
	params.Add("scope", scope)

	authUrl.RawQuery = params.Encode()
	
	http.Redirect(response, request, authUrl.String(), http.StatusTemporaryRedirect)
}