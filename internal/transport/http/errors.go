package http

import "net/http"

func internalServerError(w http.ResponseWriter, err error) {
	writeJSONError(w, err.Error(), http.StatusInternalServerError)
}

func badRequest(w http.ResponseWriter, err error) {
	writeJSONError(w, err.Error(), http.StatusBadRequest)
}
