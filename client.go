package bithavocid

import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
    "strings"
    "encoding/json"
)

type ClientBase struct {
    client *http.Client
}

type Client interface {
    LogIn(credentials Credentials) (code AuthorizationCode, err error)
}

func NewClient(appId string) Client {
    c := &ClientBase {
        client: &http.Client{},
    }
    return c
}

func (client *ClientBase)perform(path string, form url.Values, resultObject interface{}) (resp *http.Response, err error) {
    req, err := http.NewRequest("POST", fmt.Sprintf("%s/apps/%s/%s", baseURL, appId, path), strings.NewReader(form.Encode()))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("X-BITHAVOC-REQUEST-TYPE", "API")
    resp, err = client.client.Do(req)
    body := resp.Body
    defer body.Close()
    resultData, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(resultData, &resultObject)
    if err != nil {
        return nil, err
    }
    return resp, err
}
