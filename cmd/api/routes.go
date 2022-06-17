package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/h4ckm03d/simpleplan/router"
)

type logRequest struct {
	Method   string `json:"method,omitempty"`
	Path     string `json:"path,omitempty"`
	Duration int64  `json:"duration"`
	Err      string `json:"error,omitempty"`
}

// Middleware to set content type
func restMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set header
		w.Header().Set("Content-Type", "application/json")

		// Continue flow
		next.ServeHTTP(w, r)

	})
}

func errHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := f(w, r)
		var errMessage string

		if err != nil {
			errMessage = err.Error()
		}

		json.NewEncoder(log.Writer()).Encode(logRequest{
			Method:   r.Method,
			Path:     r.URL.Path,
			Duration: time.Since(start).Milliseconds(),
			Err:      errMessage,
		})
	}
}

func (app *application) routes() router.Router {
	// Create route
	r := router.New("/v1")
	r.Wrap(restMiddleware)
	r.Add(http.MethodGet, "/health", errHandler(app.healthcheckHandler))
	r.Add(http.MethodGet, "/plan", errHandler(app.getAllPlanHandler))
	r.Add(http.MethodGet, "/plan/:id", errHandler(app.getPlanHandler))

	return r
}
