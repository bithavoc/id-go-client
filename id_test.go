package bithavocid

import (
	"net/url"
	"testing"
)

func TestSuccessLogin(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "bill@gates.com")
	form.Set("password", "msdos")
	form.Set("password_confirmation", "msdos")
	resultObject := &logInResult{}
	resultObject.AuthCode = "validCode"

	twc.subscribe("myAppId", "sign-in", form, resultObject, nil)

	client := newClient("myAppId", twc)
	authCode, err := client.LogIn(Credentials{
		Email:    "bill@gates.com",
		Password: "msdos",
	})
	assertf(t, err == nil, "err should be null")
	assertf(t, authCode.Code == "validCode", "authCode should have the valid code returned by the API")
}
