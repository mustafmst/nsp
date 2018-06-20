package config

import (
	"encoding/json"
	"net/http"

	"github.com/mustafmst/nsp/core"
)

// BasicBuilder - builds app config
type BasicBuilder struct {
}

func (b BasicBuilder) ConfigRoutes(l core.Logger) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		l.LogInfo("Get -> /")
		json.NewEncoder(w).Encode("Hello world!")
	})
}

// NewBasicBuilder - creates new builder instance
func NewBasicBuilder() BasicBuilder {
	return BasicBuilder{}
}
