package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	})
}
