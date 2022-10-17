package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type KaryawanRepository interface {
	Insert(newKaryawan *model.Karyawan) error
	List(page int, totalRows int) ([]model.Karyawan, error)
	Update(newKaryawan *model.Karyawan) error
	Get(id string) (model.Karyawan, error)
	Delete(id string) error
}

type karyawanRepository struct {
	db *sqlx.DB
}

func (k *karyawanRepository) Insert(newKaryawan *model.Karyawan) error {
	_, err := k.db.NamedExec(utils.InsertKaryawan, newKaryawan)
	if err != nil {
		return err
	}
	return nil
}

func (k *karyawanRepository) List(page int, totalRows int) ([]model.Karyawan, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var karyawan []model.Karyawan
	err := k.db.Select(&karyawan, utils.SelectAllKaryawan, limit, offset)
	if err != nil {
		return nil, err
	}
	return karyawan, nil
}

func (k *karyawanRepository) Get(id string) (model.Karyawan, error) {
	var karyawan model.Karyawan
	err := k.db.Get(&karyawan, utils.SelectKaryawanById, id)
	if err != nil {
		return model.Karyawan{}, err
	}
	return karyawan, nil
}

func (k *karyawanRepository) Update(newKaryawan *model.Karyawan) error {
	_, err := k.db.NamedExec(utils.UpdateKaryawan, newKaryawan)
	if err != nil {
		return err
	}
	return nil
}

func (k *karyawanRepository) Delete(id string) error {
	_, err := k.db.Exec(utils.DeleteKaryawan, id)
	if err != nil {
		return err
	}
	return nil
}

func NewKaryawanRepository(db *sqlx.DB) KaryawanRepository {
	repo := new(karyawanRepository)
	repo.db = db
	return repo
}
