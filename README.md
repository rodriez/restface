# Restface

Restface is a helper library for make rest services

## Installation

```bash
go get github.com/rodriez/restface
```

## Usage

```go
import "github.com/rodriez/restface"

func Ping(res http.ResponseWriter, req *http.Request) {
    presenter := restface.Presenter{Writer: res}

    body := make(map[string]bool)
	body["pong"] = true

    presenter.Present(body)
}

func Error(res http.ResponseWriter, req *http.Request) {
    presenter := restface.Presenter{Writer: res}

    err := restface.NotFound("Resource not found")

    presenter.PresentError(err)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
            validUser := func(name, pass string) bool {
                return name == "Jhon" && pass == "VeryHardPass"
            }
			authenticator := &restface.NewBasicAuthenticator(req, validUser)

			if err := authenticator.Authenticate(); err != nil {
				presenter := restface.Presenter{Writer: res}
				presenter.PresentError(restface.Forbidden())

				return
			}

			next.ServeHTTP(res, req)
		})
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

[Hey 👋 buy me a beer! ](https://www.buymeacoffee.com/rodriez)