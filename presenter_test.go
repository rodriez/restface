package restface_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rodriez/restface"
)

func TestPresenter_Present(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := restface.NewPresenter(recorder)

	body := map[string]int{"id": 1234}

	err := presenter.Present(http.StatusOK, body)

	if err != nil {
		t.Errorf("no error was expected but get %s", err.Error())
	}

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code 200 but recived %d", recorder.Code)
	}

	if recorder.Body == nil {
		t.Error("Expected body with content but recived nil")
	}
}

func TestPresenter_PresentError(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := restface.NewPresenter(recorder)

	badrequest := restface.BadRequest("Invalid something")
	err := presenter.PresentError(badrequest)

	if err != nil {
		t.Errorf("no error was expected but get %s", err.Error())
	}

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400 but recived %d", recorder.Code)
	}

	if recorder.Body == nil {
		t.Error("Expected body with content but recived nil")
	}
}

func TestPresenter_PresentResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := restface.NewPresenter(recorder)

	body := strings.NewReader(`{"id":2134}`)
	response := http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(body),
		Header:     http.Header{},
	}

	response.Header.Add("Content-type", "application/json")

	err := presenter.PresentResponse(&response)

	if err != nil {
		t.Errorf("no error was expected but get %s", err.Error())
	}

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code 201 but recived %d", recorder.Code)
	}

	if recorder.Body == nil {
		t.Error("Expected body with content but recived nil")
	}
}

type mockReader struct{}

func (reader *mockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Oops something went wrong")
}

func TestPresenter_PresentResponse_ReaderError(t *testing.T) {
	recorder := httptest.NewRecorder()
	presenter := restface.NewPresenter(recorder)

	body := &mockReader{}
	response := http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(body),
		Header:     http.Header{},
	}

	response.Header.Add("Content-type", "application/json")

	err := presenter.PresentResponse(&response)

	if err == nil {
		t.Error("an error was expected but get nil")
	}
}
