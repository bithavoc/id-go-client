package bithavocid

import (
	"net/url"
)

type confirmResult struct {
	baseResult
	AuthCode string `json:"auth_code"`
}

func (client *ClientBase) Confirm(code string) (AuthorizationCode, error) {
	form := url.Values{}
	form.Set("code", code)

	resultObject := confirmResult{}
	if err := client.perform("confirm", form, &resultObject); err != nil {
		return AuthorizationCode{}, err
	}
	if err := resultObject.checkErrors(); err != nil {
		return AuthorizationCode{}, err
	}
	return AuthorizationCode{
		Code: resultObject.AuthCode,
	}, nil
}
