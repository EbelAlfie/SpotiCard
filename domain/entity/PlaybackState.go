package entity

type PlaybackStateResponse struct {
	PlayerState PlayerStateResponse `json:"player_state"`
}

type PlayerStateResponse struct {
	Track CurrentTrack `json:"track"`
}

type CurrentTrack struct {
	Uri string `json:"uri"`
	Uid string `json:"uid"`
}
