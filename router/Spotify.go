package router

import (
	"net/http"

	"spoti-card.com/controller"
)

func SpotifyRoute(server *http.ServeMux) {
	server.HandleFunc("/", controller.SpotifyController)
}
