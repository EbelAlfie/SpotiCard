package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"spoti-card.com/domain/entity"
	"spoti-card.com/domain/usecase"
)

type TrackRepositoryImpl struct {
}

func InitSpotifyRepository() usecase.TrackRepository {
	return &TrackRepositoryImpl{}
}

func (repo *TrackRepositoryImpl) GetDeviceState() {

}

func (repo *TrackRepositoryImpl) GetTrackById(trackId string) (*entity.TrackResponse, error) {
	trackUrl := fmt.Sprintf(
		"https://api.spotify.com/v1/tracks?ids=%s&market=from_token",
		trackId,
	)

	client := http.Client{}

	request, err := http.NewRequest("GET", trackUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("authorization", `Bearer ${this.authorization}`)
	request.Header.Add("client-token", `Bearer ${this.authorization}`)
	request.Header.Add("Referer", "https://open.spotify.com/")
	request.Header.Add("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	request.Header.Add("sec-ch-ua-platform", "Linux")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	request.Header.Add("sec-ch-ua-mobile", "?0")

	response, httpErr := client.Do(request)
	if httpErr != nil {
		return nil, httpErr
	}
	defer response.Body.Close()

	var result entity.TrackResponse

	decodeErr := json.NewDecoder(response.Body).Decode(&result)
	if decodeErr != nil {
		return nil, decodeErr
	}

	return &result, nil
}
