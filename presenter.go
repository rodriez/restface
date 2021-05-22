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

func (p *Presenter) Present(body interface{}) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(http.StatusOK)

	json.NewEncoder(p.Writer).Encode(body)
}

func (p *Presenter) PresentError(e PresentableError) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(e.Code())
	json.NewEncoder(p.Writer).Encode(e)
}
