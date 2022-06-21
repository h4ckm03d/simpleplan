package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/h4ckm03d/simpleplan/model"
	"github.com/h4ckm03d/simpleplan/port"
	"github.com/h4ckm03d/simpleplan/repo"
	"github.com/h4ckm03d/simpleplan/router"
	"github.com/stretchr/testify/assert"
)

type compare struct {
	want   any
	data   any
	status int
	seed   []*model.Plan
}

type testTime struct {
	port.TimeProvider
}

func (t *testTime) Now() time.Time {
	now, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	return now
}

func Test_main(t *testing.T) {

	customTime := &testTime{}

	tests := map[string]compare{
		"GET /v1/health": {
			want:   map[string]string{"status": "ok", "env": "test", "version": version},
			status: http.StatusOK,
		},
		"GET /v1/plan/1": {
			want:   model.Plan{ID: 1, Name: "Test plan", CreatedAt: customTime.Now(), UpdatedAt: customTime.Now()},
			seed:   []*model.Plan{{Name: "Test plan"}},
			status: http.StatusOK,
		},
		"GET /v1/plan": {
			want:   []model.Plan{{ID: 1, Name: "Test plan", CreatedAt: customTime.Now(), UpdatedAt: customTime.Now()}},
			seed:   []*model.Plan{{Name: "Test plan"}},
			status: http.StatusOK,
		},
		"POST /v1/plan": {
			want:   model.Plan{ID: 2, Name: "Test plan 2", CreatedAt: customTime.Now(), UpdatedAt: customTime.Now()},
			seed:   []*model.Plan{{Name: "Test plan"}},
			data:   model.Plan{Name: "Test plan 2"},
			status: http.StatusCreated,
		},
		"PUT /v1/plan/1": {
			want:   model.Plan{ID: 1, Name: "Test plan 2", CreatedAt: customTime.Now(), UpdatedAt: customTime.Now()},
			seed:   []*model.Plan{{Name: "Test plan"}},
			data:   model.Plan{ID: 1, Name: "Test plan 2", CreatedAt: customTime.Now(), UpdatedAt: customTime.Now()},
			status: http.StatusOK,
		},
		"POST /v1/plan empty": {
			want:   nil,
			seed:   []*model.Plan{{Name: "Test plan"}},
			data:   nil,
			status: http.StatusBadRequest,
		},
		"DELETE /v1/plan/1": {
			want:   nil,
			seed:   []*model.Plan{{Name: "Test plan"}},
			data:   nil,
			status: http.StatusOK,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			requests := strings.Split(name, " ")
			if len(requests) < 2 {
				t.Fatal("invalid test name")
			}
			buff := new(bytes.Buffer)
			if tt.data != nil {
				json.NewEncoder(buff).Encode(tt.data)
			}
			// Create a request to pass to our handler.
			req, err := http.NewRequest(requests[0], requests[1], buff)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

			// Create a new instance of the application.
			app := &application{
				config:   config{env: "test"},
				logger:   logger,
				PlanRepo: repo.NewPlanRepo(customTime),
			}

			for _, seed := range tt.seed {
				app.PlanRepo.Create(seed)
			}

			wantBuffer := new(bytes.Buffer)
			err = json.NewEncoder(wantBuffer).Encode(tt.want)
			assert.Nil(t, err)

			router.Build(app.routes()).ServeHTTP(rr, req)
			assert.Equal(t, tt.status, rr.Result().StatusCode)
			if tt.status == http.StatusOK && tt.want != nil {
				assert.Equal(t, wantBuffer.String(), rr.Body.String())
			}
		})
	}
}
