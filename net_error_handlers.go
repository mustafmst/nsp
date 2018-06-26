package nsp

import (
	"fmt"
	"net/http"
)

func noSuchMethodInPath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "No %s method defined under %s path", r.Method, r.RequestURI)
}
