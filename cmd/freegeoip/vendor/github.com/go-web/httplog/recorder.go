package httplog

import "net/http"

// The ResponseRecorder interface is implemented by ResponseWriters that
// can record the response status code and bytes written to the client.
type ResponseRecorder interface {
	Code() int  // Response status code
	Bytes() int // Bytes written to the client
}

// ResponseWriter is an http.ResponseWriter + ResponseRecorder.
type ResponseWriter struct {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier
	ResponseRecorder
	bytes, code int
}

// NewResponseWriter creates and initializes a new ResponseWriter.
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		code:           http.StatusOK,
	}
}

// Header implements the http.ResponseWriter interface.
func (w *ResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Write implements the http.ResponseWriter interface.
func (w *ResponseWriter) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

// WriteHeader implements the http.ResponseWriter interface.
func (w *ResponseWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

// Code implements the ResponseRecorder interface.
func (w *ResponseWriter) Code() int { return w.code }

// Bytes implements the ResponseRecorder interface.
func (w *ResponseWriter) Bytes() int { return w.bytes }
