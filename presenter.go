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

//Present - Write a json body
func (p *Presenter) Present(body interface{}) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(http.StatusOK)

	json.NewEncoder(p.Writer).Encode(body)
}

//PresentError - Write a json error
func (p *Presenter) PresentError(e PresentableError) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(e.Code())
	json.NewEncoder(p.Writer).Encode(e)
}
