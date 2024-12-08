package data

import "spoti-card.com/domain/usecase"

type TokenRepositoryImpl struct{}

func TokenRepository() usecase.TokenRepository {
	return &TokenRepositoryImpl{}
}

func (repo *TokenRepositoryImpl) FetchRefreshToken() {

}

func (repo *TokenRepositoryImpl) FetchClientToken() {

}
