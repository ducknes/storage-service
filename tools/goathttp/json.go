package goathttp

import (
	"encoding/json"
	"net/http"
)

func ReadRequestJson(r *http.Request, v any) error {
	if v == nil {
		return nil
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteResponseJson(w http.ResponseWriter, statusCode int, v any) error {
	if v == nil {
		return nil
	}

	w.Header().Set(_contentTypeHeader, _contentTypeJSON)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
