package entity

type TrackResponse struct {
	Tracks []TrackEntity `json:"tracks"`
}

type TrackEntity struct {
	Album       AlbumEntity    `json:"album"`
	Artists     []ArtistEntity `json:"artists"`
	DiscNumber  int            `json:"disc_number"`
	DurationMs  int64          `json:"duration_ms"`
	Explicit    bool           `json:"explicit"`
	Href        string         `json:"href"`
	Id          string         `json:"id"`
	IsLocal     string         `json:"is_local"`
	IsPlayable  bool           `json:"is_playable"`
	Name        string         `json:"name"`
	Popularity  int            `json:"popularity"`
	PreviewUrl  string         `json:"preview_url"`
	TrackNumber string         `json:"track_number"`
	Type        string         `json:"type"`
	Uri         string         `json:"uri"`
}

type AlbumEntity struct {
	Album_Type string  `json:"album_type"`
	Images     []Image `json:"images"`
}

type ArtistEntity struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
