package repo

import (
	"sort"
	"sync"
	"time"

	"github.com/h4ckm03d/simpleplan/model"
	"github.com/h4ckm03d/simpleplan/port"
)

type PlanRepo struct {
	Id     int
	Data   map[int]*model.Plan
	ListId []int
	m      sync.Mutex
}

var _ port.PlanRepo = &PlanRepo{}

func NewPlanRepo() *PlanRepo {
	return &PlanRepo{
		Id:     0,
		Data:   make(map[int]*model.Plan),
		ListId: []int{},
	}
}
func (r *PlanRepo) Create(plan *model.Plan) (*model.Plan, error) {
	r.m.Lock()
	defer r.m.Unlock()
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()
	r.Id++
	plan.ID = r.Id
	r.Data[plan.ID] = plan
	r.ListId = append(r.ListId, plan.ID)
	return plan, nil
}

func (r *PlanRepo) Get(id int) (*model.Plan, error) {
	r.m.Lock()
	defer r.m.Unlock()
	plan, ok := r.Data[int(id)]
	if !ok {
		return nil, model.ErrNotFound
	}

	return plan, nil
}

func (r *PlanRepo) Update(plan *model.Plan) (*model.Plan, error) {
	r.m.Lock()
	defer r.m.Unlock()
	_, found := r.Data[plan.ID]
	if plan.ID == 0 || plan.ID > r.Id || !found {
		return nil, model.ErrNotFound
	}

	plan.UpdatedAt = time.Now()
	r.Data[plan.ID] = plan
	return plan, nil
}

func (r *PlanRepo) Delete(id int) error {
	r.m.Lock()
	defer r.m.Unlock()
	if _, found := r.Data[id]; !found {
		return model.ErrNotFound
	}

	delete(r.Data, int(id))
	// data always sorted because listId is incremental id
	index := sort.SearchInts(r.ListId, id)
	r.ListId = append(r.ListId[:index], r.ListId[index+1:]...)
	return nil
}

func (r *PlanRepo) GetAll(limit, page int) ([]*model.Plan, error) {
	r.m.Lock()
	defer r.m.Unlock()
	plans := make([]*model.Plan, 0)
	totalLen := len(r.ListId)

	start := page * limit
	end := start + limit
	if start >= totalLen {
		return plans, nil
	}

	if end > totalLen {
		end = totalLen
	}

	for ; start < end && start < totalLen; start++ {
		if plan, ok := r.Data[r.ListId[start]]; ok {
			plans = append(plans, plan)
		}
	}

	return plans, nil
}
