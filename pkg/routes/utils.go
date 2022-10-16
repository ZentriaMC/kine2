package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendJSON(w http.ResponseWriter, code int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(response)
}

func sendResponse(w http.ResponseWriter, response interface{}) error {
	return sendJSON(w, http.StatusOK, response)
}

func sendError(w http.ResponseWriter, code int, response interface{}) error {
	return sendJSON(w, code, struct {
		Message interface{} `json:"message"`
	}{response})
}

func parseJSON(r *http.Request, w http.ResponseWriter, data interface{}) (res bool) {
	ctype := r.Header.Get("Content-Type")
	if ctype != "application/json" {
		sendError(w, http.StatusUnsupportedMediaType, fmt.Sprintf("Bad Content-Type %s, accept application/json", ctype))
		return
	}

	dec := json.NewDecoder(r.Body)
	//dec.DisallowUnknownFields()

	if err := dec.Decode(data); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	return true
}
