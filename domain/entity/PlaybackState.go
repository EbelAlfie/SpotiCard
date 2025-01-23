package entity

type PlayerStateResponse struct {
	Track        TrackEntity  `json:"item"`
	Progress int       `json:"progress_ms"`
	IsPlaying    bool         `json:"is_playing"`
}
