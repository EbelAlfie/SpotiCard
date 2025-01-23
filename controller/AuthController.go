package controller

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func AuthController(response http.ResponseWriter, request *http.Request) {
	rawUrl := "https://accounts.spotify.com/authorize?"
	
	authUrl, err := url.Parse(rawUrl)
	if err != nil { //Sebenernya ga perlu
		log.Fatal("error parsing url")
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	clientID := os.Getenv("CLIENT_ID")
	redirectTarg := "http://localhost:3031"
	scope := "user-read-currently-playing"

	params := authUrl.Query()
	params.Add("client_id", clientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirectTarg)
	params.Add("scope", scope)

	authUrl.RawQuery = params.Encode()
	
	http.Redirect(response, request, authUrl.String(), http.StatusTemporaryRedirect)
}