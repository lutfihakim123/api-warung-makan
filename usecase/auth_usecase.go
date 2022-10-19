package usecase

import (
	"api-warung-makan/model"
	"api-warung-makan/repository"
)

type AuthUseCase interface {
	LoginKaryawan(username string, password string) (model.Credential, error)
}

type authUseCase struct {
	repo repository.AuthRepository
}

func (a *authUseCase) LoginKaryawan(username string, password string) (model.Credential, error) {
	return a.repo.Login(username, password)
}

func NewAuthUseCase(repo repository.AuthRepository) AuthUseCase {
	pc := new(authUseCase)
	pc.repo = repo
	return pc
}
