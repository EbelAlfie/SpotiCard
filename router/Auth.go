package router

import (
	"net/http"

	"spoti-card.com/controller"
)

func AuthRouter(server *http.ServeMux) {
	server.HandleFunc("/login", controller.AuthController)
}