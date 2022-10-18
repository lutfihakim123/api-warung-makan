package repository

import (
	"api-warung-makan/model"
	"api-warung-makan/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type NotaRepository interface {
	Insert(newNota *model.Nota) error
	List(page int, totalRows int) ([]model.Nota, error)
	Update(newNota *model.Nota) error
	Get(id string) (model.Nota, error)
	Delete(id string) error
}

type notaRepository struct {
	db *sqlx.DB
}

func (n *notaRepository) Insert(newNota *model.Nota) error {
	_, err := n.db.NamedExec(utils.InsertNota, newNota)
	if err != nil {
		return err
	}
	fmt.Println("sukses")
	return nil
}

func (n *notaRepository) List(page int, totalRows int) ([]model.Nota, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var nota []model.Nota
	err := n.db.Select(&nota, utils.SelectAllNota, limit, offset)
	if err != nil {
		return nil, err
	}
	return nota, nil
}

func (n *notaRepository) Get(id string) (model.Nota, error) {
	var nota model.Nota
	err := n.db.Get(&nota, utils.SelectNotaById, id)
	if err != nil {
		return model.Nota{}, err
	}
	return nota, nil
}

func (n *notaRepository) Update(newNota *model.Nota) error {
	_, err := n.db.NamedExec(utils.UpdateNota, newNota)
	if err != nil {
		return err
	}
	return nil
}

func (n *notaRepository) Delete(id string) error {
	_, err := n.db.Exec(utils.DeleteNota, id)
	if err != nil {
		return err
	}
	return nil
}

func NewNotaRepository(db *sqlx.DB) NotaRepository {
	repo := new(notaRepository)
	repo.db = db
	return repo
}
