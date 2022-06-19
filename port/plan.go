package port

import "github.com/h4ckm03d/simpleplan/model"

type PlanRepo interface {
	Create(plan *model.Plan) (*model.Plan, error)
	Get(id int) (*model.Plan, error)
	Update(plan *model.Plan) (*model.Plan, error)
	Delete(id int) error
	GetAll(limit, page int) ([]*model.Plan, error)
}
