package restface

type Authenticator interface {
	Authenticate() error
}
