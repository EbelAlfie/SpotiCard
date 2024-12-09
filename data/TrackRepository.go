package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"spoti-card.com/domain/entity"
	"spoti-card.com/domain/usecase"
)

type TrackRepositoryImpl struct {
	accessToken string
	clientToken string
}

func TrackRepository(
	accessToken entity.AccessTokenEntity,
	clientToken entity.ClientTokenEntity,
) usecase.TrackRepository {
	return &TrackRepositoryImpl{
		accessToken: fmt.Sprintf("Bearer %s", accessToken.AccessToken),
		clientToken: clientToken.GrantedToken.Token,
	}
}

func (repo *TrackRepositoryImpl) GetDeviceState() (*entity.PlaybackStateResponse, error) {
	url := "https://gew4-spclient.spotify.com/connect-state/v1/devices/hobs_86133792d6f7240c655de45fa6bc7f30527"

	client := http.Client{}

	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	_, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *TrackRepositoryImpl) GetTrackById(trackId string) (*entity.TrackEntity, error) {
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
	request.Header.Add("client-token", repo.accessToken)
	request.Header.Add("Referer", "https://open.spotify.com/")
	request.Header.Add("sec-ch-ua", `"Microsoft Edge";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
	request.Header.Add("sec-ch-ua-platform", "Linux")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")
	request.Header.Add("sec-ch-ua-mobile", "?0")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result entity.TrackResponse

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Tracks[0], nil
}
