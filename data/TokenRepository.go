package data

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"spoti-card.com/domain/entity"
	"spoti-card.com/domain/usecase"
)

type TokenRepositoryImpl struct{
	code string
}

func TokenRepository(code string) usecase.TokenRepository {
	return &TokenRepositoryImpl{
		code: code,
	}
}

func (repo *TokenRepositoryImpl) FetchAccessToken() (*entity.TokenResponse, error) {
	uri := "https://accounts.spotify.com/api/token"

	clientSecret := os.Getenv("CLIENT_SECRET")
	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" || clientSecret == "" {
		return nil, fmt.Errorf("no secret or id provided")
	}

	auth := base64.StdEncoding.
		EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret))) 

	body := url.Values {
		"grant_type": {"authorization_code"},
		"code": {repo.code},
		"redirect_uri": {"http://localhost:3031"},
	}

	request, err := http.NewRequest("POST", uri, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Basic " + auth)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, &entity.HttpError { 
			StatusCode: response.StatusCode,
			Message: response.Status,
		}
	}
	
	var result *entity.TokenResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}