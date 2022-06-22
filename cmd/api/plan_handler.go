package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/h4ckm03d/simpleplan/model"
	"github.com/h4ckm03d/simpleplan/router"
)

func (app *application) planHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return app.getAllPlanHandler(w, r)
	case "POST":
		return app.createPlanHandler(w, r)
	case "DELETE":
	}

	return errors.New("not found")
}

func (app *application) planMutationHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return app.getPlanHandler(w, r)
	case "PUT":
		return app.updatePlanHandler(w, r)
	case "DELETE":
		return app.deletePlanHandler(w, r)
	}

	return errors.New("not found")
}

func (app *application) getPlanHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(router.Param(r, "id"))
	if err != nil {
		return err
	}

	data, err := app.PlanRepo.Get(id)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(data)
}

func (app *application) updatePlanHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(router.Param(r, "id"))
	if err != nil {
		return err
	}
	var update *model.Plan
	defer dclose(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return err
	}

	update.ID = id

	data, err := app.PlanRepo.Update(update)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(data)
}

func (app *application) createPlanHandler(w http.ResponseWriter, r *http.Request) error {
	if r.ContentLength == 0 {
		return errors.New("empty body")
	}

	var plan *model.Plan
	defer dclose(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		return err
	}

	data, err := app.PlanRepo.Create(plan)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(data)
}

func (app *application) getAllPlanHandler(w http.ResponseWriter, r *http.Request) error {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if limit == 0 || limit > 100 {
		limit = 10
	}

	data, err := app.PlanRepo.GetAll(limit, page)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(data)
}

func (app *application) deletePlanHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(router.Param(r, "id"))
	if err != nil {
		return err
	}

	return app.PlanRepo.Delete(id)
}
