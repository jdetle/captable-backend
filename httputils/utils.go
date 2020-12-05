// Package httputils provides http utilities.
package httputils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// ValidateJSONPayload tries to JSON Unmarshal body into payload (which is
// passed as a pointer). It will return the first improper type in the error.
func ValidateJSONPayload(w http.ResponseWriter, body io.ReadCloser, payload interface{}) error {
	badErr := errors.New("bad payload")
	err := json.NewDecoder(body).Decode(payload)
	if err != nil {
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			HTTPError(w,
				fmt.Sprintf("wrong type for %s, expected %s, got %s",
					typeErr.Field, typeErr.Type, typeErr.Value),
				http.StatusBadRequest, err)
			return badErr
		}
		HTTPError(w, "unable to parse json body", http.StatusBadRequest, err)
		return badErr
	}
	return nil
}

// HTTPError returns a error message to a client.
func HTTPError(w http.ResponseWriter, msg string, status int, err error) {
	if err != nil {
		log.Errorf("%s: %#v", msg, err)
	}
	message := ErrHTTPResponse{Message: msg, Code: status}
	SendJSON(w, status, message)
}

// SendJSON encodes thing as json and returns it to w with status.
func SendJSON(w http.ResponseWriter, status int, thing interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	err := json.NewEncoder(w).Encode(thing)
	if err != nil {
		// something is wacky
		log.Errorf("error returning response: %#v", err)
		http.Error(w, "error returning response", http.StatusInternalServerError)
	}
}

// ErrHTTPResponse couples an http response code with an error message.
type ErrHTTPResponse struct {
	Message string // error message string
	Code    int    // status code
}

func (e *ErrHTTPResponse) Error() string {
	return e.Message
}

// Status returns the HTTP Status.
func (e *ErrHTTPResponse) Status() int {
	return e.Code
}

// NewErrHTTPResponse is a functional creator for an ErrHTTPResponse struct.
func NewErrHTTPResponse(m string, c int) error {
	return &ErrHTTPResponse{
		Message: m,
		Code:    c,
	}
}
