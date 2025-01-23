package usecase

import "spoti-card.com/domain/entity"

type TrackRepository interface {
	GetPlaybackState() (*entity.PlayerStateResponse, error)

	GetTrackById(trackId string) (*entity.TrackEntity, error)
}