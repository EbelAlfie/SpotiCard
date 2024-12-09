package entity

type PlaybackStateResponse struct {
	Track CurrentTrack `json:"track"`
}

type CurrentTrack struct {
	Uri string `json:"uri"`
	Uid string `json:"uid"`
}
