package restface

import (
	"encoding/json"
	"net/http"
)

type Presenter struct {
	Writer http.ResponseWriter
}

func (p *Presenter) Present(body interface{}) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(http.StatusOK)

	json.NewEncoder(p.Writer).Encode(body)
}

func (p *Presenter) PresentError(err *ApiError) {
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(err.StatusCode)
	json.NewEncoder(p.Writer).Encode(err)
}