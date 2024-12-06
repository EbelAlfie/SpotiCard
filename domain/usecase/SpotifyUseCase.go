package usecase

type SpotifyRepository interface {
	GetSpotifyCard() string

	AuthenticateSpotify()

	ConnectState()
}
