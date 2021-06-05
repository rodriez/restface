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
    presenter := restface.NewPresenter(res)

    body := make(map[string]bool)
	body["pong"] = true

    presenter.Present(http.StatusOK, body)
}

func Error(res http.ResponseWriter, req *http.Request) {
    presenter := restface.NewPresenter(res)

    err := restface.NotFound("Resource not found")

    presenter.PresentError(err)
}

func AuthMiddleware(next http.Handler) http.Handler {
	validUser := func(name, pass string) bool {
		return name == "Jhon" && pass == "VeryHardPass"
	}
	authenticator := restface.NewBasicAuthenticator(validUser)

	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
            if err := authenticator.Authenticate(req); err != nil {
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