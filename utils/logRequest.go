package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	log.Printf("Request: %s %s", r.Method, r.URL.String())

	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("Header: %s: %s", name, value)
		}
	}

	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			log.Printf("Body: %s", string(body))
			r.Body = io.NopCloser(bytes.NewReader(body))
		} else {
			log.Printf("Error reading body: %v", err)
		}
	}
}
