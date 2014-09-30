package bithavocid

import (
	"net/url"
)

func (client *ClientBase) Forgot(code, password string) error {
	form := url.Values{}
	form.Set("code", code)
	form.Set("password", password)
	form.Set("password_confirmation", password)

	resultObject := baseResult{}
	if err := client.perform("forgot", form, &resultObject); err != nil {
		return err
	}
	if err := resultObject.checkErrors(); err != nil {
		return err
	}
    return nil
}
