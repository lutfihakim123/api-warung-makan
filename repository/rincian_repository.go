package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type RincianRepository interface {
	Insert(newRincian *model.Rincian) error
	List(page int, totalRows int) ([]model.Rincian, error)
	Update(newRincian *model.Rincian) error
	Get(id string) (model.Rincian, error)
	Delete(id string) error
	GetMenu(id string) (model.Menu, error)
}

type rincianRepository struct {
	db *sqlx.DB
}

func (r *rincianRepository) Insert(newRincian *model.Rincian) error {
	tx := r.db.MustBegin()
	tx.NamedExec(utils.UpdateStockMenu, &newRincian)
	tx.NamedExec(utils.InsertRincian, newRincian)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *rincianRepository) List(page int, totalRows int) ([]model.Rincian, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var rincian []model.Rincian
	err := r.db.Select(&rincian, utils.SelectAllRincian, limit, offset)
	if err != nil {
		return nil, err
	}
	return rincian, nil
}

func (r *rincianRepository) Get(id string) (model.Rincian, error) {
	var rincian model.Rincian
	err := r.db.Get(&rincian, utils.SelectRincianById, id)
	if err != nil {
		return model.Rincian{}, err
	}
	return rincian, nil
}

func (r *rincianRepository) Update(newRincian *model.Rincian) error {
	tx := r.db.MustBegin()
	tx.NamedExec(utils.UpdateStockMenu, &newRincian)
	tx.NamedExec(utils.UpdateRincian, newRincian)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *rincianRepository) Delete(id string) error {
	_, err := r.db.Exec(utils.DeleteRincian, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *rincianRepository) GetMenu(id string) (model.Menu, error) {
	var menu model.Menu
	err := r.db.Get(&menu, utils.SelectMenuById, id)
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func NewRincianRepository(db *sqlx.DB) RincianRepository {
	repo := new(rincianRepository)
	repo.db = db
	return repo
}
