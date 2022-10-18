package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type MejaRepository interface {
	Insert(newMeja *model.Meja) error
	List(page int, totalRows int) ([]model.Meja, error)
	Update(newMeja *model.Meja) error
	Get(id string) (model.Meja, error)
	Delete(id string) error
}

type mejaRepository struct {
	db *sqlx.DB
}

func (m *mejaRepository) Insert(newMeja *model.Meja) error {
	_, err := m.db.NamedExec(utils.InsertMeja, newMeja)
	if err != nil {
		return err
	}
	return nil
}

func (m *mejaRepository) List(page int, totalRows int) ([]model.Meja, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var meja []model.Meja
	err := m.db.Select(&meja, utils.SelectAllMeja, limit, offset)
	if err != nil {
		return nil, err
	}
	return meja, nil
}

func (m *mejaRepository) Get(id string) (model.Meja, error) {
	var meja model.Meja
	err := m.db.Get(&meja, utils.SelectMejaById, id)
	if err != nil {
		return model.Meja{}, err
	}
	return meja, nil
}

func (m *mejaRepository) Update(newMeja *model.Meja) error {
	_, err := m.db.NamedExec(utils.UpdateMeja, newMeja)
	if err != nil {
		return err
	}
	return nil
}

func (m *mejaRepository) Delete(id string) error {
	_, err := m.db.Exec(utils.DeleteMeja, id)
	if err != nil {
		return err
	}
	return nil
}

func NewMejaRepository(db *sqlx.DB) MejaRepository {
	repo := new(mejaRepository)
	repo.db = db
	return repo
}
