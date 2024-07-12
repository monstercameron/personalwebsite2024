package routes

import (
    "bytes"
    "net/http"
)

// statusRecorder wraps the http.ResponseWriter to capture the status code, header state, and response body
type statusRecorder struct {
    http.ResponseWriter
    status      int
    wroteHeader bool
    body        *bytes.Buffer
}

// WriteHeader captures the status code and sets the header written flag
func (rec *statusRecorder) WriteHeader(status int) {
	if rec.wroteHeader {
		return
	}
	rec.status = status
	rec.wroteHeader = true
}

// Write writes data to the body buffer and sets the header if not already written
func (rec *statusRecorder) Write(b []byte) (int, error) {
	if !rec.wroteHeader {
		rec.WriteHeader(http.StatusOK)
	}
	return rec.body.Write(b)
}
