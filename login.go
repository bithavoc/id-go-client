package bithavocid

import (
	"net/url"
)

type UserInfo struct {
	Fullname string
	Email    string
}

type User struct {
	Token Token
	Info  UserInfo
}

type Token struct {
	Code string
}

type Credentials struct {
	Email    string
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
	form := url.Values{}
	form.Set("email", info.Email)
	form.Set("password", info.Password)
	form.Set("password_confirmation", info.Password)

	if err != nil {
		return AuthorizationCode{}, err
	}

	resultObject := logInResult{}
	err = client.perform("sign-in", form, &resultObject)
	if err != nil {
		return AuthorizationCode{}, err
	}
	err = resultObject.checkErrors()
	if err != nil {
		return AuthorizationCode{}, err
	}
	authCode := AuthorizationCode{
		Code: resultObject.AuthCode,
	}
	return authCode, nil
}

type negotiateResult struct {
	baseResult
	Token string
	User  UserInfo
}

func (client *ClientBase) Negotiate(code AuthorizationCode) (User, error) {
	form := url.Values{}
	form.Set("code", code.Code)

	resultObject := negotiateResult{}
	err := client.perform("tokens", form, &resultObject)
	if err != nil {
		return User{}, err
	}
	err = resultObject.checkErrors()
	if err != nil {
		return User{}, err
	}
	user := User{
		Info: resultObject.User,
		Token: Token{
			Code: resultObject.Token,
		},
	}
	return user, nil
}
