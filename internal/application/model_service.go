package application

import "github.com/AllanXavierRA/vectrag/internal/domain"

type ModelService struct {
	repo domain.Repository
}

func NewModelService(repo domain.Repository) *ModelService {
	return &ModelService{
		repo: repo,
	}
}

func (ms *ModelService) Create(model domain.Model) error {
	return ms.repo.CreateModel(model)
}

func (ms *ModelService) Update(model domain.Model) error {
	return ms.repo.UpdateModel(model)
}

func (ms *ModelService) Delete(id string) error {
	return ms.repo.DeleteModel(id)
}

func (ms *ModelService) Get(id string) (domain.Model, error) {
	return ms.repo.GetModel(id)
}

func (ms *ModelService) List() ([]domain.Model, error) {
	return ms.repo.GetModels()
}
