package bithavocid

import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "strings"
    "encoding/json"
    "os"
)

const (
    appId = "59abb3124156a6e47e39108e36f9f380"
)

func getBaseURL() string {
    if os.Getenv("BH_ENV") != "" {
        return "http://127.0.0.1:4000"
    }
    return "https://id.bithavoc.io"
}

type ClientBase struct {
    client *http.Client
}

type Client interface {
    LogIn(credentials Credentials) (code AuthorizationCode, err error)
    SignUp(info SignUp) (err error)
    Negotiate(code AuthorizationCode) (user User, err error)
}

func NewClient(appId string) Client {
    c := &ClientBase {
        client: &http.Client{},
    }
    return c
}

func (client *ClientBase)perform(path string, form url.Values, resultObject interface{}) (resp *http.Response, err error) {
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps/%s/%s", getBaseURL(), appId, path), strings.NewReader(form.Encode()))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("X-BITHAVOC-REQUEST-TYPE", "API")
    resp, err = client.client.Do(req)
    if err != nil {
        return nil, err
    }
    body := resp.Body
    defer func() {
        if body != nil {
            body.Close()
        }
    }()
    resultData, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(resultData, &resultObject)
    if err != nil {
        return nil, err
    }
    return resp, err
}
