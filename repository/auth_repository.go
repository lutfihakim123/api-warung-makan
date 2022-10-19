package repository

import (
	"api-warung-makan/model"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Login(username string, password string) (model.Credential, error)
}

type authRepository struct {
	db *sqlx.DB
}

func (a *authRepository) Login(username string, password string) (model.Credential, error) {
	var credential model.Credential
	err := a.db.Get(&credential, "select id, nama, alamat, gaji, username, password from mst_karyawan where username=$1 and password=$2", username, password)
	if err != nil {
		return model.Credential{}, err
	}
	return credential, nil
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	repo := new(authRepository)
	repo.db = db
	return repo
}
