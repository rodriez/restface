package restface_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rodriez/restface"
)

func TestAuthenticate_NoError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost", strings.NewReader(`{"id":1234}`))
	req.SetBasicAuth("john", "pass")

	userExist := func(name, pass string) bool {
		return true
	}

	authenticator := restface.NewBasicAuthenticator(userExist)
	if err := authenticator.Authenticate(req); err != nil {
		t.Errorf("no error was expected but get %s", err.Error())
	}
}

func TestAuthenticate_Error(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost", strings.NewReader(`{"id":1234}`))
	userExist := func(name, pass string) bool {
		return false
	}

	authenticator := restface.NewBasicAuthenticator(userExist)
	if err := authenticator.Authenticate(req); err == nil {
		t.Error("an error was expected but get nil")
	}
}
