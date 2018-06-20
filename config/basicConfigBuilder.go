package config

import (
	"encoding/json"
	"net/http"
)

// BasicBuilder - builds app config
type BasicBuilder struct {
}

func (b BasicBuilder) ConfigRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world!")
	})
}

// NewBasicBuilder - creates new builder instance
func NewBasicBuilder() BasicBuilder {
	return BasicBuilder{}
}
