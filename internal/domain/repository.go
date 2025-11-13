package domain

type Repository interface {
	CreateModel(model Model) error
	UpdateModel(model Model) error
	DeleteModel(id string) error
	GetModel(id string) (Model, error)
	GetModels() ([]Model, error)
}
