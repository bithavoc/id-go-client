package bithavocid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getBaseURL() string {
	if os.Getenv("BH_ENV") != "" {
		return "http://127.0.0.1:4000"
	}
	return "https://id.bithavoc.io"
}

type ClientBase struct {
	client webClient
	appId  string
}

func (client *ClientBase) perform(path string, form url.Values, resultObject interface{}) (err error) {
	return client.client.perform(client.GetAppId(), path, form, resultObject)
}

type Client interface {
	LogIn(credentials Credentials) (code AuthorizationCode, err error)
	SignUp(info SignUp) (err error)
	Negotiate(code AuthorizationCode) (user User, err error)
	Confirm(code string) (AuthorizationCode, error)
	Recover(email string) error
	SetAppId(appId string)
	GetAppId() string
}

func NewClient(appId string) Client {
	return newClient(appId, newLiveWebClient())
}

func newClient(appId string, wc webClient) Client {
	c := &ClientBase{
		client: wc,
	}
	c.SetAppId(appId)
	return c
}

func (client *ClientBase) GetAppId() string {
	return client.appId
}

func (client *ClientBase) SetAppId(appId string) {
	client.appId = appId
}

type webClient interface {
	perform(appId string, path string, form url.Values, resultObject interface{}) error
}

type liveWebClient struct {
	client *http.Client
}

func newLiveWebClient() webClient {
	return &liveWebClient{
		client: &http.Client{},
	}
}

func (client *liveWebClient) perform(appId string, path string, form url.Values, resultObject interface{}) (err error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps/%s/%s", getBaseURL(), appId, path), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-BITHAVOC-REQUEST-TYPE", "API")
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	body := resp.Body
	defer func() {
		if body != nil {
			body.Close()
		}
	}()
	resultData, err := ioutil.ReadAll(resp.Body)
	//    fmt.Println(string(resultData))
	err = json.Unmarshal(resultData, &resultObject)
	if err != nil {
		return err
	}
	return err
}
