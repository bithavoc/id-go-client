package bithavocid

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

type testWebCall struct {
	requestId    string
	path         string
	appId        string
	form         url.Values
	resultObject interface{}
	err          error
}

func (call *testWebCall) generateRequestId() {
	call.requestId = fmt.Sprintf("%s|%s|%s", call.appId, call.path, call.form.Encode())
}

type testWebClient struct {
	requests map[string]*testWebCall
}

func newTestWebClient() *testWebClient {
	c := &testWebClient{
		requests: make(map[string]*testWebCall),
	}
	return c
}

func (client *testWebClient) subscribe(appId string, path string, form url.Values, resultObject interface{}, err error) {
	call := &testWebCall{
		appId:        appId,
		path:         path,
		form:         form,
		resultObject: resultObject,
		err:          err,
	}
	call.generateRequestId()
	_, ok := client.requests[call.requestId]
	if ok {
		panic(fmt.Errorf("Repeated subscription to %s %s", appId, path))
	} else {
		client.requests[call.requestId] = call
	}
}

func (client *testWebClient) perform(appId string, path string, form url.Values, resultObject interface{}) error {
	call := &testWebCall{
		appId: appId,
		path:  path,
		form:  form,
	}
	call.generateRequestId()
	mockedCall, ok := client.requests[call.requestId]
	if ok {
		outParam := reflect.ValueOf(resultObject).Elem()
		mockedResult := reflect.ValueOf(mockedCall.resultObject).Elem()
		outParam.Set(mockedResult)
		return mockedCall.err
	} else {
		panic(fmt.Errorf("Mocked call not found for %s %s", appId, path))
	}
}

func assertf(t *testing.T, condition bool, format string, args ...interface{}) {
	if !condition {
		t.Errorf(format, args...)
	}
}
