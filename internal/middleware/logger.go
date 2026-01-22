package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RequestLog struct {
	RequestID string    `json:"request_id"`
	Agent     string    `json:"agent"`
	URL       string    `json:"url"`
	Method    string    `json:"method"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

// RequestLogger middleware untuk logging HTTP request
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate request ID
		requestID := uuid.New().String()

		var body string
		if r.Body != nil && (r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH") {
			bodyBytes, err := io.ReadAll(r.Body)
			if err == nil {
				body = string(bodyBytes)
				// Restore body untuk handler selanjutnya
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// Get User-Agent
		agent := r.UserAgent()
		if agent == "" {
			agent = "Unknown"
		}

		// Create request log
		requestLog := RequestLog{
			RequestID: requestID,
			Agent:     agent,
			URL:       r.URL.String(),
			Method:    r.Method,
			Body:      body,
			Timestamp: time.Now(),
		}

		// Convert to JSON and log
		logJSON, err := json.Marshal(requestLog)
		if err != nil {
			log.Printf("Error marshaling request log: %v", err)
		} else {
			log.Printf("%s", string(logJSON))
		}

		// Add request ID to response header (optional, untuk tracing)
		w.Header().Set("X-Request-ID", requestID)

		// Call next handler
		next.ServeHTTP(w, r)
	})
}
