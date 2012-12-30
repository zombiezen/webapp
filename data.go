package webapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Common HTTP headers
const (
	HeaderAllow              = "Allow"
	HeaderContentDisposition = "Content-Disposition"
	HeaderContentEncoding    = "Content-Encoding"
	HeaderContentLength      = "Content-Length"
	HeaderContentType        = "Content-Type"
)

// Content types
const (
	HTMLType = "text/html; charset=utf-8"
	JSONType = "application/json; charset=utf-8"
)

// MethodNotAllowed replies to a request with an HTTP 405 method not allowed error.
func MethodNotAllowed(w http.ResponseWriter, methods ...string) {
	w.Header().Set(HeaderAllow, strings.Join(methods, ", "))
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// ContentLength sets the Content-Length header to size.
func ContentLength(h http.Header, size int64) {
	h.Set(HeaderContentLength, strconv.FormatInt(size, 10))
}

// Attachment sets the Content-Disposition header to an attachment with the given file name.
func Attachment(h http.Header, base, ext string) {
	h.Set(HeaderContentDisposition, "attachment; filename="+base+"."+ext)
}

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
