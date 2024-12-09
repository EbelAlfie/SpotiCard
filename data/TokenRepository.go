package data

import (
	"encoding/json"
	"net/http"

	"spoti-card.com/domain/entity"
	"spoti-card.com/domain/usecase"
)

type TokenRepositoryImpl struct{}

func TokenRepository() usecase.TokenRepository {
	return &TokenRepositoryImpl{}
}

func (repo *TokenRepositoryImpl) FetchAccessToken() (*entity.AccessTokenEntity, error) {
	url := "https://open.spotify.com/get_access_token?reason=transport&productType=web-player"

	client := http.Client{}

	request, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Add("cookie", `sp_dc=${this.me};`)

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	var result *entity.AccessTokenEntity

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *TokenRepositoryImpl) FetchClientToken() (*entity.ClientTokenEntity, error) {
	url := "https://clienttoken.spotify.com/v1/clienttoken"

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	var responseResult *entity.ClientTokenEntity

	err = json.NewDecoder(response.Body).Decode(&responseResult)
	if err != nil {
		return nil, err
	}

	return responseResult, nil
}
