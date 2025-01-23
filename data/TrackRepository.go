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
}

func TrackRepository(
	accessToken string,
) usecase.TrackRepository {
	return &TrackRepositoryImpl{
		accessToken: fmt.Sprintf("Bearer %s", accessToken),
	}
}

func (repo *TrackRepositoryImpl) GetPlaybackState() (*entity.PlayerStateResponse, error) {
	url := "https://api.spotify.com/v1/me/player/currently-playing"
	
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", repo.accessToken)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result *entity.PlayerStateResponse
	err = json.NewDecoder(response.Body).Decode(&result) 
	if err != nil {
		return nil, err
	}

	return result, nil
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

	request.Header.Add("Authorization", repo.accessToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result *entity.TrackResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Tracks[0], nil
}
