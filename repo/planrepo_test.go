package store_test

import (
	"testing"

	"github.com/h4ckm03d/simpleplan/model"
	"github.com/h4ckm03d/simpleplan/store"
	"github.com/stretchr/testify/assert"
)

func TestPlanRepo_Create(t *testing.T) {
	// Create plan repo
	r := store.NewPlanRepo()

	// Create plan
	plan := &model.Plan{
		Name: "Test plan",
	}

	// Create plan
	plan, err := r.Create(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")
}

func TestPlanRepo_Get(t *testing.T) {
	// Create plan repo
	r := store.NewPlanRepo()

	// Create plan
	plan := &model.Plan{
		Name: "Test plan",
	}

	// Create plan
	plan, err := r.Create(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")

	// Get plan
	plan, err = r.Get(plan.ID)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")
}

func TestPlanRepo_Update(t *testing.T) {
	// Create plan repo
	r := store.NewPlanRepo()

	// Create plan
	plan := &model.Plan{
		Name: "Test plan",
	}

	// Create plan
	plan, err := r.Create(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")

	// Update plan
	plan.Name = "Test plan updated"
	plan, err = r.Update(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan updated")

	unknown := &model.Plan{
		ID:   1000,
		Name: "Test plan",
	}

	plan, err = r.Update(unknown)
	assert.Error(t, err)
	assert.Nil(t, plan)
}

func TestPlanRepo_Delete(t *testing.T) {

	// Create plan repo
	r := store.NewPlanRepo()

	// Create plan
	plan := &model.Plan{
		Name: "Test plan",
	}

	// Create plan
	plan, err := r.Create(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")

	// Delete plan
	err = r.Delete(plan.ID)
	assert.NoError(t, err)

	// Delete unknown
	err = r.Delete(1000)
	assert.Error(t, err)

	// Get plan
	plan, err = r.Get(plan.ID)
	assert.Error(t, err)
	assert.Nil(t, plan)
}

func TestPlanRepo_GetAll(t *testing.T) {

	// Create plan repo
	r := store.NewPlanRepo()

	// Create plan
	plan := &model.Plan{
		Name: "Test plan",
	}

	// Create plan
	plan, err := r.Create(plan)
	assert.NoError(t, err)
	assert.NotNil(t, plan)
	assert.NotNil(t, plan.ID)
	assert.NotNil(t, plan.CreatedAt)
	assert.NotNil(t, plan.UpdatedAt)
	assert.Equal(t, plan.Name, "Test plan")

	// Get all plans
	plans, err := r.GetAll(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, plans)
	assert.Equal(t, len(plans), 1)
	assert.Equal(t, plans[0].Name, "Test plan")

	// Get all but empty results
	plans, err = r.GetAll(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, []*model.Plan{}, plans)
}
