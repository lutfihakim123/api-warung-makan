package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
)

type KaryawanUseCase interface {
	CreateNewKaryawan(newKaryawan *model.Karyawan) error
	GetAllKaryawan(page int, totalRows int) ([]model.Karyawan, error)
	UpdateKaryawan(newKaryawan model.Karyawan) error
	DeleteKaryawan(id string) error
	GetKaryawanById(id string) (model.Karyawan, error)
}

type karyawanUseCase struct {
	repo repository.KaryawanRepository
}

func (k *karyawanUseCase) CreateNewKaryawan(newKaryawan *model.Karyawan) error {
	newKaryawan.Id = utils.GenerateId()
	return k.repo.Insert(newKaryawan)
}

func (k *karyawanUseCase) GetAllKaryawan(page int, totalRows int) ([]model.Karyawan, error) {
	return k.repo.List(page, totalRows)
}

func (k *karyawanUseCase) UpdateKaryawan(newKaryawan model.Karyawan) error {
	return k.repo.Update(&newKaryawan)
}

func (k *karyawanUseCase) DeleteKaryawan(id string) error {
	return k.repo.Delete(id)
}

func (k *karyawanUseCase) GetKaryawanById(id string) (model.Karyawan, error) {
	return k.repo.Get(id)
}

func NewKaryawanUseCase(repo repository.KaryawanRepository) KaryawanUseCase {
	pc := new(karyawanUseCase)
	pc.repo = repo
	return pc
}
