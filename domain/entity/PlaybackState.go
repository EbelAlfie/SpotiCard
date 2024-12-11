package entity

type PlaybackStateResponse struct {
	PlayerState PlayerStateResponse `json:"player_state"`
}

type PlayerStateResponse struct {
	Track        CurrentTrack `json:"track"`
	LastPosition string       `json:"position_as_of_timestamp"`
	Duration     string       `json:"duration"`
	IsPlaying    bool         `json:"is_playing"`
	IsPaused     bool         `json:"is_paused"`
}

type CurrentTrack struct {
	Uri string `json:"uri"`
	Uid string `json:"uid"`
}
