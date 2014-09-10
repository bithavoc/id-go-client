package bithavocid

import (
	"net/url"
)

type SignUp struct {
	Email                string
	Password             string
	PasswordConfirmation string
	Fullname             string
}

func (client *ClientBase) SignUp(info SignUp) (err error) {
	form := url.Values{}
	form.Set("email", info.Email)
	form.Set("fullname", info.Fullname)
	form.Set("password", info.Password)
	form.Set("password_confirmation", info.Password)

	resultObject := baseResult{}
	if err = client.perform("sign-up", form, &resultObject); err != nil {
		return err
	}
	if err = resultObject.checkErrors(); err != nil {
		return err
	}
	return nil
}
