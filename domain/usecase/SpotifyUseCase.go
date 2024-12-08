package usecase

import "spoti-card.com/domain/entity"

type TrackRepository interface {
	GetDeviceState()

	GetTrackById(trackId string) (*entity.TrackResponse, error)
}

type TokenRepository interface {
	FetchRefreshToken()

	FetchClientToken()
}
