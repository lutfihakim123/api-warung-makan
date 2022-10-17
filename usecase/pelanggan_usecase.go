package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
)

type PelangganUseCase interface {
	CreateNewPelanggan(newPelanggan *model.Pelanggan) error
	GetAllPelanggan(page int, totalRows int) ([]model.Pelanggan, error)
	UpdatePelanggan(newPelanggan model.Pelanggan) error
	DeletePelanggan(id string) error
	GetPelangganById(id string) (model.Pelanggan, error)
}

type pelangganUseCase struct {
	repo repository.PelangganRepository
}

func (p *pelangganUseCase) CreateNewPelanggan(newPelanggan *model.Pelanggan) error {
	newPelanggan.Id = utils.GenerateId()
	return p.repo.Insert(newPelanggan)
}

func (p *pelangganUseCase) GetAllPelanggan(page int, totalRows int) ([]model.Pelanggan, error) {
	return p.repo.List(page, totalRows)
}

func (p *pelangganUseCase) UpdatePelanggan(newPelanggan model.Pelanggan) error {
	return p.repo.Update(&newPelanggan)
}

func (p *pelangganUseCase) DeletePelanggan(id string) error {
	return p.repo.Delete(id)
}

func (p *pelangganUseCase) GetPelangganById(id string) (model.Pelanggan, error) {
	return p.repo.Get(id)
}

func NewPelangganUseCase(repo repository.PelangganRepository) PelangganUseCase {
	pc := new(pelangganUseCase)
	pc.repo = repo
	return pc
}
