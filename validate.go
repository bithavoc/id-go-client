package bithavocid

type BasicUserInfo struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type tokenValidation struct {
	baseResult
	User BasicUserInfo `json:"user"`
}

func (client *ClientBase) Validate(token string) (BasicUserInfo, error) {
	resultObject := tokenValidation{}
	if err := client.retrieve("tokens/"+token, &resultObject); err != nil {
		return resultObject.User, err
	}
	if err := resultObject.checkErrors(); err != nil {
		return resultObject.User, err
	}
	return resultObject.User, nil
}
