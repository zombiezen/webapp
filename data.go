package webapp

import (
	"encoding/json"
	"net/http"
)

// Content types
const (
	HTMLType = "text/html; charset=utf-8"
	JSONType = "application/json; charset=utf-8"
)

// JSONResponse writes a JSON value to w, setting the Content-Type.
func JSONResponse(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", JSONType)
	return json.NewEncoder(w).Encode(v)
}
