package bithavocid

import (
	"fmt"
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

func TestFailedLogin(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "bill@gates.com")
	form.Set("password", "obviously_wrong")
	form.Set("password_confirmation", "obviously_wrong")
	resultObject := &logInResult{}
	twc.subscribe("myAppId", "sign-in", form, resultObject, fmt.Errorf("password: password is incorrect"))

	client := newClient("myAppId", twc)
	authCode, err := client.LogIn(Credentials{
		Email:    "bill@gates.com",
		Password: "obviously_wrong",
	})
	assertf(t, err != nil, "error should not be null")
	assertf(t, err.Error() == "password: password is incorrect", "password error should be given")
	assertf(t, authCode.Code == "", "authCode should be empty")
}
