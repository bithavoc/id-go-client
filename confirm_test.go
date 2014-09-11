package bithavocid

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSuccessConfirm(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "bill@gates.com")
	form.Set("fullname", "Bill G")
	form.Set("password", "msdos")
	form.Set("password_confirmation", "msdos")
	resultObject := &baseResult{}

	twc.subscribe("myAppId", "sign-up", form, resultObject, nil)

	client := newClient("myAppId", twc)
	err := client.SignUp(SignUp{
		Email:    "bill@gates.com",
		Password: "msdos",
		Fullname: "Bill G",
	})
	assertf(t, err == nil, "err should be null")
}

func TestFailedConfirm(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "bill@gates.com")
	form.Set("fullname", "")
	form.Set("password", "msdos")
	form.Set("password_confirmation", "msdos")
	resultObject := &baseResult{}
	twc.subscribe("myAppId", "sign-up", form, resultObject, fmt.Errorf("fullname: fullname is required"))

	client := newClient("myAppId", twc)
	err := client.SignUp(SignUp{
		Email:    "bill@gates.com",
		Password: "msdos",
	})
	assertf(t, err != nil, "error should not be null")
	assertf(t, err.Error() == "fullname: fullname is required", "value required error should be given")
}
