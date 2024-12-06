package data

import "spoti-card.com/domain/usecase"

type SpotifyRepositoryImpl struct {
}

func InitSpotifyRepository() usecase.SpotifyRepository {
	return &SpotifyRepositoryImpl{}
}

func (repo *SpotifyRepositoryImpl) GetSpotifyCard() string {
	return `
		<h1>HAHAY</h1>
	`
}

func (repo *SpotifyRepositoryImpl) AuthenticateSpotify() {

}

func (repo *SpotifyRepositoryImpl) ConnectState() {

}
