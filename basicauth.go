package restface

import (
	"errors"
	"net/http"
)

type basicAuthenticator struct {
	request   *http.Request
	userExist func(name, pass string) bool
}

func (a *basicAuthenticator) Authenticate() error {
	user, pass, ok := a.request.BasicAuth()

	if ok && a.userExist(user, pass) {
		return nil
	}

	return errors.New("forbidden")
}

func NewBasicAuthenticator(req *http.Request, userExist func(name, pass string) bool) Authenticator {
	return &basicAuthenticator{
		request:   req,
		userExist: userExist,
	}
}
