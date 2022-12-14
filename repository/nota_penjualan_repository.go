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
	Get(id string) (model.Nota, error)
	ListReport(page int, totalRows int) ([]model.ReportPenjualan, error)
	GetReport(id string) (model.ReportPenjualan, error)
	Update(newNota *model.Nota) error
	GetMeja(id string) (model.Meja, error)
	Delete(id string) error
}

type notaRepository struct {
	db *sqlx.DB
}

func (n *notaRepository) Insert(newNota *model.Nota) error {
	tx := n.db.MustBegin()
	tx.NamedExec(utils.UpdateStatusMeja, newNota)
	tx.NamedExec(utils.InsertNota, newNota)
	err := tx.Commit()
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

func (n *notaRepository) ListReport(page int, totalRows int) ([]model.ReportPenjualan, error) {
	limit := totalRows
	offset := limit * (page - 1)
	var report []model.ReportPenjualan
	err := n.db.Select(&report, utils.SelectAllReport, limit, offset)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (n *notaRepository) Get(id string) (model.Nota, error) {
	var nota model.Nota
	err := n.db.Get(&nota, utils.SelectNotaById, id)
	if err != nil {
		return model.Nota{}, err
	}
	return nota, nil
}

func (n *notaRepository) GetReport(id string) (model.ReportPenjualan, error) {
	var report model.ReportPenjualan
	err := n.db.Get(&report, utils.SelectReportById, id)
	if err != nil {
		return model.ReportPenjualan{}, err
	}
	return report, nil
}

func (n *notaRepository) GetMeja(id string) (model.Meja, error) {
	var meja model.Meja
	err := n.db.Get(&meja, utils.SelectMejaById, id)
	if err != nil {
		return model.Meja{}, err
	}
	return meja, nil
}

func (n *notaRepository) Update(newNota *model.Nota) error {
	tx := n.db.MustBegin()
	tx.NamedExec(utils.UpdateStatusMeja, newNota)
	tx.NamedExec(utils.UpdateNota, newNota)
	err := tx.Commit()
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
