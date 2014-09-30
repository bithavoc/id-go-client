package bithavocid

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSuccessForgot(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("code", "emailedCode")
	form.Set("password", "MyNewPassw0rd")
	form.Set("password_confirmation", "MyNewPassw0rd")
	resultObject := &baseResult{}

	twc.subscribe("myAppId", "forgot", form, resultObject, nil)

	client := newClient("myAppId", twc)
	err := client.Forgot("emailedCode", "MyNewPassw0rd")
	assertf(t, err == nil, "err should be null")
}

func TestFailedForgot(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("code", "invalidCode")
	form.Set("password", "MyNewPassw0rd")
	form.Set("password_confirmation", "MyNewPassw0rd")
	resultObject := &baseResult{}
	twc.subscribe("myAppId", "forgot", form, resultObject, fmt.Errorf("The given code is invalid"))

	client := newClient("myAppId", twc)
	err := client.Forgot("invalidCode", "MyNewPassw0rd")
	assertf(t, err != nil, "error should not be null")
	assertf(t, err.Error() == "The given code is invalid", "value required error should be given")
}
