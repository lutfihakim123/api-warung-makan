package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
	"api-warung-makan/utils"
)

type MenuUseCase interface {
	CreateNewMenu(newMenu *model.Menu) error
	GetAllMenu(page int, totalRows int) ([]model.Menu, error)
	UpdateMenu(newMenu model.Menu) error
	DeleteMenu(id string) error
	GetMenuById(id string) (model.Menu, error)
}

type menuUseCase struct {
	repo repository.MenuRepository
}

func (m *menuUseCase) CreateNewMenu(newMenu *model.Menu) error {
	newMenu.Id = utils.GenerateId()
	return m.repo.Insert(newMenu)
}

func (m *menuUseCase) GetAllMenu(page int, totalRows int) ([]model.Menu, error) {
	return m.repo.List(page, totalRows)
}

func (m *menuUseCase) UpdateMenu(newMenu model.Menu) error {
	return m.repo.Update(&newMenu)
}

func (m *menuUseCase) DeleteMenu(id string) error {
	return m.repo.Delete(id)
}

func (m *menuUseCase) GetMenuById(id string) (model.Menu, error) {
	return m.repo.Get(id)
}

func NewMenuUseCase(repo repository.MenuRepository) MenuUseCase {
	pc := new(menuUseCase)
	pc.repo = repo
	return pc
}
