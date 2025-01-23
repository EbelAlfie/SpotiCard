package utils

import (
	"net/http"

	"spoti-card.com/domain/entity"
	"spoti-card.com/presentation"
)

func HandleError(response http.ResponseWriter, err error, statusCode int) {
	status := statusCode
	if instance, res := err.(*entity.HttpError); res {
		status = instance.StatusCode
	}

	errorModel := presentation.ErrorModel{
		Error: err,
	}

	errorCard := presentation.ErrorCard(errorModel)

	response.WriteHeader(status)
	response.Write([]byte(errorCard))
}
