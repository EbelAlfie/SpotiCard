package entity

type AccessTokenEntity struct {
	ClientId                         string `json:"client_id"`
	AccessToken                      string `json:"accessToken"`
	AccessTokenExpirationTimestampMs int64  `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous                      bool   `json:"isAnonymous"`
}

type ClientTokenEntity struct {
	ResponseType string             `json:"response_type"`
	GrantedToken GrantedTokenEntity `json:"granted_token"`
}

type GrantedTokenEntity struct {
	Token   string         `json:"token"`
	Expire  int64          `json:"expires_after_seconds"`
	Refresh int64          `json:"refresh_after_seconds"`
	Domains []DomainEntity `json:"domains"`
}

type DomainEntity struct {
	Domain string `json:"domain"`
}
