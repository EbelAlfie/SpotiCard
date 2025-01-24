package usecase

import "spoti-card.com/domain/entity"

type TrackRepository interface {
	GetPlaybackState() (*entity.PlayerStateResponse, error)

	GetRecentlyPlayed() (*entity.TrackEntity, error)

	GetTrackById(trackId string) (*entity.TrackEntity, error)
}