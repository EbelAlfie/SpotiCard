package domain

type SpotifyRepository interface {
	GetSpotifyCard() string

	AuthenticateSpotify()
}
