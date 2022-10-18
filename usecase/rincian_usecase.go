package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
)

type RincianUseCase interface {
	CreateNewRincian(newRincian *model.Rincian) error
	GetAllRincian(page int, totalRows int) ([]model.Rincian, error)
	UpdateRincian(newRincian model.Rincian) error
	DeleteRincian(id string) error
	GetRincianById(id string) (model.Rincian, error)
	GetMenuById(id string) (model.Menu, error)
}

type rincianUseCase struct {
	repo repository.RincianRepository
}

func (r *rincianUseCase) CreateNewRincian(newRincian *model.Rincian) error {
	newRincian.Id = utils.GenerateId()
	return r.repo.Insert(newRincian)
}

func (r *rincianUseCase) GetAllRincian(page int, totalRows int) ([]model.Rincian, error) {
	return r.repo.List(page, totalRows)
}

func (r *rincianUseCase) UpdateRincian(newRincian model.Rincian) error {
	return r.repo.Update(&newRincian)
}

func (r *rincianUseCase) DeleteRincian(id string) error {
	return r.repo.Delete(id)
}

func (r *rincianUseCase) GetRincianById(id string) (model.Rincian, error) {
	return r.repo.Get(id)
}

func (r *rincianUseCase) GetMenuById(id string) (model.Menu, error) {
	return r.repo.GetMenu(id)
}

func NewRincianUseCase(repo repository.RincianRepository) RincianUseCase {
	pc := new(rincianUseCase)
	pc.repo = repo
	return pc
}
