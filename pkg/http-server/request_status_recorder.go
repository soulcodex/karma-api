package httpserver

import "net/http"

type RequestStatusRecorder struct {
	http.ResponseWriter
	Status int
}

func NewRequestStatusRecorder(w http.ResponseWriter) *RequestStatusRecorder {
	return &RequestStatusRecorder{ResponseWriter: w, Status: http.StatusOK}
}

func (rsr *RequestStatusRecorder) WriteHeader(status int) {
	rsr.Status = status
	rsr.ResponseWriter.WriteHeader(status)
}
