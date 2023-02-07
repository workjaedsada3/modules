package jwt

import "github.com/workjaedsada3/modules/helpers"

type (
	Token struct {
		AccessToken      string      `json:"access_token"`
		ExpiresIn        interface{} `json:"expires_in"`
		RefreshExpiresIn interface{} `json:"refresh_expires_in"`
		RefreshToken     string      `json:"refresh_token"`
		TokenType        string      `json:"token_type"`
	}
	TokenDetail struct {
		Id         *helpers.UUID `json:"id"`
		Firstname  string        `json:"first_name"`
		MiddleName string        `json:"middlename"`
		Lastname   string        `json:"last_name"`
		Email      string        `json:"email"`
	}
	RefreshToken struct {
		RefreshExpiresIn interface{} `json:"refresh_expires_in"`
		RefreshToken     string      `json:"refresh_token"`
	}
)
