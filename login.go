package bithavocid

import (
    "net/url"
)

type Token struct {
    Code string
}

type Credentials struct {
    Email string
    Password string
}

type AuthorizationCode struct {
    Code string
}

type baseResult struct {
    Messages MappedErrorList
}

type logInResult struct {
    baseResult
    AuthCode string `json:"auth_code"`
}

func (client *ClientBase) LogIn(info Credentials) (code AuthorizationCode, err error) {
    form :=  url.Values{}
    form.Set("email", info.Email)
    form.Set("password", info.Password)
    form.Set("password_confirmation", info.Password)

    if err != nil {
        return AuthorizationCode{}, err
    }

    resultObject := logInResult{}
    _, err = client.perform("sign-up", form, &resultObject)
    if err != nil {
        return AuthorizationCode{}, err
    }
    err = resultObject.checkErrors()
    if err != nil {
        return AuthorizationCode{}, err
    }
    authCode := AuthorizationCode {
        Code: resultObject.AuthCode,
    }
    return authCode, nil
}
