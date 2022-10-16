package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"

	"github.com/jmoiron/sqlx"
)

type MenuRepository interface {
	Insert(newMenu *model.Menu) error
	List(page int, totalRows int) ([]model.Menu, error)
	Update(newMenu *model.Menu) error
	Get(id string) (model.Menu, error)
	Delete(id string) error
}

type menuRepository struct {
	db *sqlx.DB
}

func (m *menuRepository) Insert(newMenu *model.Menu) error {
	_, err := m.db.NamedExec(utils.InsertMenu, newMenu)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuRepository) List(page int, totalRows int) ([]model.Menu, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var menu []model.Menu
	err := m.db.Select(&menu, utils.SelectAllMenu, limit, offset)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (m *menuRepository) Get(id string) (model.Menu, error) {
	var menu model.Menu
	err := m.db.Get(&menu, utils.SelectMenuById, id)
	if err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (m *menuRepository) Update(newMenu *model.Menu) error {
	_, err := m.db.NamedExec(utils.UpdateMenu, newMenu)
	if err != nil {
		return err
	}
	return nil
}

func (m *menuRepository) Delete(id string) error {
	_, err := m.db.Exec(utils.DeleteMenu, id)
	if err != nil {
		return err
	}
	return nil
}

func NewMenuRepository(db *sqlx.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
