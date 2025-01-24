package entity

type RecentlyPlayedResponse struct {
	Tracks []TrackModel `json:"items"`
}

type TrackModel struct {
	Track TrackEntity `json:"track"`
}