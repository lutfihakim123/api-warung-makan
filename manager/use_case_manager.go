package manager

import "api-warung-makan/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
	KaryawanUseCase() usecase.KaryawanUseCase
	PelangganUseCase() usecase.PelangganUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (u *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(u.repo.MenuRepository())
}

func (u *useCaseManager) KaryawanUseCase() usecase.KaryawanUseCase {
	return usecase.NewKaryawanUseCase(u.repo.KaryawanRepository())
}

func (u *useCaseManager) PelangganUseCase() usecase.PelangganUseCase {
	return usecase.NewPelangganUseCase(u.repo.PelangganRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
