package utils

import (
	"net/http"

	"spoti-card.com/presentation"
)

func HandleError(response http.ResponseWriter, err error, statusCode int) {
	var errorModel = presentation.ErrorModel{
		Error: err,
	}

	errorCard := presentation.ErrorCard(errorModel)

	response.WriteHeader(statusCode)
	response.Header().Set("Content-Type", "image/svg+xml")
	response.Write([]byte(errorCard))
}
