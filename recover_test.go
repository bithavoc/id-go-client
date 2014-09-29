package bithavocid

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSuccessRecover(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "bill@gates.com")
	resultObject := &baseResult{}

	twc.subscribe("myAppId", "recover", form, resultObject, nil)

	client := newClient("myAppId", twc)
	err := client.Recover("bill@gates.com")
	assertf(t, err == nil, "err should be null")
}

func TestFailedRecover(t *testing.T) {
	twc := newTestWebClient()
	form := url.Values{}
	form.Set("email", "doesntExist@gates.com")
	resultObject := &baseResult{}
	twc.subscribe("myAppId", "recover", form, resultObject, fmt.Errorf("No user found with the given email address"))

	client := newClient("myAppId", twc)
	err := client.Recover("doesntExist@gates.com")
	assertf(t, err != nil, "error should not be null")
	assertf(t, err.Error() == "No user found with the given email address", "value required error should be given")
}
