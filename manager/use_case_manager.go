package manager

import "api-warung-makan/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
	KaryawanUseCase() usecase.KaryawanUseCase
	PelangganUseCase() usecase.PelangganUseCase
	NotaUseCase() usecase.NotaUseCase
	MejaUseCase() usecase.MejaUseCase
	RincianUseCase() usecase.RincianUseCase
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

func (u *useCaseManager) NotaUseCase() usecase.NotaUseCase {
	return usecase.NewNotaUseCase(u.repo.NotaRepository())
}

func (u *useCaseManager) MejaUseCase() usecase.MejaUseCase {
	return usecase.NewMejaUseCase(u.repo.MejaRepository())
}

func (u *useCaseManager) RincianUseCase() usecase.RincianUseCase {
	return usecase.NewRincianUseCase(u.repo.RincianRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
