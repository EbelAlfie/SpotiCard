package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"spoti-card.com/domain/entity"
	"spoti-card.com/domain/usecase"
)

type TokenRepositoryImpl struct{}

func TokenRepository() usecase.TokenRepository {
	return &TokenRepositoryImpl{}
}

func (repo *TokenRepositoryImpl) FetchAccessToken() (*entity.AccessTokenEntity, error) {
	cookie := os.Getenv("ME")
	cookieHeader := fmt.Sprintf("sp_dc=%s;", cookie)

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

	request.Header.Add("cookie", cookieHeader)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result *entity.AccessTokenEntity

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *TokenRepositoryImpl) FetchClientToken(clientId string) (*entity.ClientTokenEntity, error) {
	url := "https://clienttoken.spotify.com/v1/clienttoken"

	client := http.Client{}

	body := fmt.Sprintf(
		`
		{
			"client_data":{
				"client_version":"1.2.53.257.g47fa6c39",
				"client_id":"%s",
				"js_sdk_data":{
					"device_brand":"unknown",
					"device_model":"unknown",
					"os":"linux",
					"os_version":"unknown",
					"device_id":"hahaha",
					"device_type":"computer"
					}
				}
			}
	`, clientId)
	reqBody := strings.NewReader(body)

	request, err := http.NewRequest("GET", url, reqBody)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseResult *entity.ClientTokenEntity

	err = json.NewDecoder(response.Body).Decode(&responseResult)
	if err != nil {
		return nil, err
	}

	return responseResult, nil
}
