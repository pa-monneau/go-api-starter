package handler

import "net/http"

type Matchers struct {
	Path    string
	Handler func(w http.ResponseWriter, _r *http.Request)
}
