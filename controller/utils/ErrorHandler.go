package utils

import (
	"net/http"

	"spoti-card.com/presentation"
)

func HandleError(err error, response http.ResponseWriter) {
	var errorModel = presentation.ErrorModel{
		Error: err,
	}
	errorCard := presentation.ErrorCard(errorModel)
	response.Write([]byte(errorCard))
}
