package utils

import (
    "log"
    "net/http"
)

// LogRequest logs the details of an incoming HTTP request.
func LogRequest(r *http.Request) {
    log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
}

// HandleError logs the error and sends an appropriate HTTP response.
func HandleError(w http.ResponseWriter, err error, statusCode int) {
    log.Printf("Error: %v", err)
    http.Error(w, http.StatusText(statusCode), statusCode)
}