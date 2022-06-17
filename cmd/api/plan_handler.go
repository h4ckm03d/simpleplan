package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/h4ckm03d/simpleplan/router"
)

func (app *application) getPlanHandler(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
		"id":      router.Param(r, "id"),
	})

	id, err := strconv.Atoi(router.Param(r, "id"))
	if err != nil {
		return err
	}

	fmt.Println("got id:", id)

	return nil
}

func (app *application) updatePlanHandler(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "update",
		"env":     app.config.env,
		"version": version,
	})

	return nil
}

func (app *application) createPlanHandler(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	})
	return nil
}

func (app *application) getAllPlanHandler(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	})
	return nil
}

func (app *application) deletePlanHandler(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	})
	return nil
}
