package data

import "spoti-card.com/domain"

type SpotifyRepositoryImpl struct {
}

func InitSpotifyRepository() domain.SpotifyRepository {
	return &SpotifyRepositoryImpl{}
}

func (repo *SpotifyRepositoryImpl) GetSpotifyCard() string {
	return `
		<h1>HAHAY</h1>
	`
}

func (repo *SpotifyRepositoryImpl) AuthenticateSpotify() {

}
