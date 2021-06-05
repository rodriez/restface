package restface

import (
	"errors"
	"net/http"
)

type basicAuthenticator struct {
	userExist func(name, pass string) bool
}

//Authenticate - Return an error is the user is not authorized
func (a *basicAuthenticator) Authenticate(request *http.Request) error {
	user, pass, ok := request.BasicAuth()

	if ok && a.userExist(user, pass) {
		return nil
	}

	return errors.New("forbidden")
}

//NewBasicAuthenticator - Return a basic authenticator
func NewBasicAuthenticator(userExist func(name, pass string) bool) Authenticator {
	return &basicAuthenticator{
		userExist: userExist,
	}
}
