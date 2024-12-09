package usecase

import "spoti-card.com/domain/entity"

type TrackRepository interface {
	GetDeviceState() (*entity.PlaybackStateResponse, error)

	GetTrackById(trackId string) (*entity.TrackEntity, error)
}

type TokenRepository interface {
	FetchAccessToken() (*entity.AccessTokenEntity, error)

	FetchClientToken(clientId string) (*entity.ClientTokenEntity, error)
}
