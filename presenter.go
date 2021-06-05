package restface

import (
	"encoding/json"
	"net/http"
)

type PresentableError interface {
	Code() int
}

type Presenter struct {
	Writer http.ResponseWriter
}

func NewPresenter(w http.ResponseWriter) *Presenter {
	return &Presenter{w}
}

//Present - Write an http response
func (p *Presenter) PresentResponse(resp *http.Response) error {
	for k, _ := range resp.Header {
		p.Writer.Header().Add(k, resp.Header.Get(k))
	}
	p.Writer.WriteHeader(resp.StatusCode)

	body := []byte{}
	if _, err := resp.Body.Read(body); err != nil {
		return err
	}

	p.Writer.Write(body)

	return nil
}

//Present - Write a 200 ok and json body
func (p *Presenter) Present(statusCode int, body interface{}) error {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(statusCode)

	return json.NewEncoder(p.Writer).Encode(body)
}

//PresentError - Write a json error
func (p *Presenter) PresentError(e PresentableError) error {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(e.Code())
	return json.NewEncoder(p.Writer).Encode(e)
}
