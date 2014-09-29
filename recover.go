package bithavocid

import (
	"net/url"
)

func (client *ClientBase) Recover(email string) error {
	form := url.Values{}
	form.Set("email", email)

	resultObject := baseResult{}
	if err := client.perform("recover", form, &resultObject); err != nil {
		return err
	}
	if err := resultObject.checkErrors(); err != nil {
		return err
	}
    return nil
}
