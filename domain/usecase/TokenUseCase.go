package usecase

import "spoti-card.com/domain/entity"

type TokenRepository interface {
	FetchAccessToken() (*entity.TokenResponse, error)
}
