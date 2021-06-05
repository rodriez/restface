package restface

import "net/http"

type Authenticator interface {
	Authenticate(req *http.Request) error
}
