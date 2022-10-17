package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type PelangganRepository interface {
	Insert(newPelanggan *model.Pelanggan) error
	List(page int, totalRows int) ([]model.Pelanggan, error)
	Update(newPelanggan *model.Pelanggan) error
	Get(id string) (model.Pelanggan, error)
	Delete(id string) error
}

type pelangganRepository struct {
	db *sqlx.DB
}

func (p *pelangganRepository) Insert(newPelanggan *model.Pelanggan) error {
	_, err := p.db.NamedExec(utils.InsertPelanggan, newPelanggan)
	if err != nil {
		return err
	}
	return nil
}

func (p *pelangganRepository) List(page int, totalRows int) ([]model.Pelanggan, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var pelanggan []model.Pelanggan
	err := p.db.Select(&pelanggan, utils.SelectAllPelanggan, limit, offset)
	if err != nil {
		return nil, err
	}
	return pelanggan, nil
}

func (p *pelangganRepository) Get(id string) (model.Pelanggan, error) {
	var pelanggan model.Pelanggan
	err := p.db.Get(&pelanggan, utils.SelectPelangganById, id)
	if err != nil {
		return model.Pelanggan{}, err
	}
	return pelanggan, nil
}

func (p *pelangganRepository) Update(newPelanggan *model.Pelanggan) error {
	_, err := p.db.NamedExec(utils.UpdatePelanggan, newPelanggan)
	if err != nil {
		return err
	}
	return nil
}

func (p *pelangganRepository) Delete(id string) error {
	_, err := p.db.Exec(utils.DeletePelanggan, id)
	if err != nil {
		return err
	}
	return nil
}

func NewPelangganRepository(db *sqlx.DB) PelangganRepository {
	repo := new(pelangganRepository)
	repo.db = db
	return repo
}
