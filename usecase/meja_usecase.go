package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
)

type MejaUseCase interface {
	CreateNewMeja(newMeja *model.Meja) error
	GetAllMeja(page int, totalRows int) ([]model.Meja, error)
	UpdateMeja(newMeja model.Meja) error
	DeleteMeja(id string) error
	GetMejaById(id string) (model.Meja, error)
}

type mejaUseCase struct {
	repo repository.MejaRepository
}

func (m *mejaUseCase) CreateNewMeja(newMeja *model.Meja) error {
	newMeja.Id = utils.GenerateId()
	return m.repo.Insert(newMeja)
}

func (m *mejaUseCase) GetAllMeja(page int, totalRows int) ([]model.Meja, error) {
	return m.repo.List(page, totalRows)
}

func (m *mejaUseCase) UpdateMeja(newMeja model.Meja) error {
	return m.repo.Update(&newMeja)
}

func (m *mejaUseCase) DeleteMeja(id string) error {
	return m.repo.Delete(id)
}

func (m *mejaUseCase) GetMejaById(id string) (model.Meja, error) {
	return m.repo.Get(id)
}

func NewMejaUseCase(repo repository.MejaRepository) MejaUseCase {
	pc := new(mejaUseCase)
	pc.repo = repo
	return pc
}
