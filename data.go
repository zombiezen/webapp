package webapp

import (
	"encoding/json"
	"fmt"
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

// A MultiError is returned by operations that have errors on particular elements.
// This is functionally identical to appengine.MultiError.
type MultiError []error

func (e MultiError) Error() string {
	msg, n := "", 0
	for _, err := range e {
		if err != nil {
			if n == 0 {
				msg = err.Error()
			}
			n++
		}
	}
	switch n {
	case 0:
		return "0 errors"
	case 1:
		return msg
	case 2:
		return msg + " (and 1 other error)"
	}
	return fmt.Sprintf("%s (and %d other errors)", msg, n-1)
}
