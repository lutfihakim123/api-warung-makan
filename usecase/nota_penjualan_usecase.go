package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
	"time"
)

type NotaUseCase interface {
	CreateNewNota(newNota *model.Nota) error
	GetAllNota(page int, totalRows int) ([]model.Nota, error)
	UpdateNota(newNota model.Nota) error
	DeleteNota(id string) error
	GetNotaById(id string) (model.Nota, error)
	GetMejaById(id string) (model.Meja, error)
}

type notaUseCase struct {
	repo repository.NotaRepository
}

func (n *notaUseCase) CreateNewNota(newNota *model.Nota) error {
	newNota.Id = utils.GenerateId()
	newNota.WaktuPesan = time.Now()
	return n.repo.Insert(newNota)
}

func (n *notaUseCase) GetAllNota(page int, totalRows int) ([]model.Nota, error) {
	return n.repo.List(page, totalRows)
}

func (n *notaUseCase) UpdateNota(newNota model.Nota) error {
	return n.repo.Update(&newNota)
}

func (n *notaUseCase) DeleteNota(id string) error {
	return n.repo.Delete(id)
}

func (n *notaUseCase) GetNotaById(id string) (model.Nota, error) {
	return n.repo.Get(id)
}
func (n *notaUseCase) GetMejaById(id string) (model.Meja, error) {
	return n.repo.GetMeja(id)
}

func NewNotaUseCase(repo repository.NotaRepository) NotaUseCase {
	pc := new(notaUseCase)
	pc.repo = repo
	return pc
}
