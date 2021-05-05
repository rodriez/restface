package restface

import (
	"errors"
	"net/http"
)

type BasicAuthenticator struct {
	Request   *http.Request
	ValidUser func(name, pass string) bool
}

func (a *BasicAuthenticator) Authenticate() error {
	user, pass, ok := a.Request.BasicAuth()

	if ok && a.ValidUser(user, pass) {
		return nil
	}

	return errors.New("forbidden")
}
