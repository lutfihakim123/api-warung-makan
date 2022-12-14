package manager

import "api-warung-makan/repository"

type RepositoryManager interface {
	MenuRepository() repository.MenuRepository
	KaryawanRepository() repository.KaryawanRepository
	PelangganRepository() repository.PelangganRepository
	NotaRepository() repository.NotaRepository
	MejaRepository() repository.MejaRepository
	RincianRepository() repository.RincianRepository
	AuthRepository() repository.AuthRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (i *repositoryManager) MenuRepository() repository.MenuRepository {
	return repository.NewMenuRepository(i.infra.SqlDB())
}

func (i *repositoryManager) KaryawanRepository() repository.KaryawanRepository {
	return repository.NewKaryawanRepository(i.infra.SqlDB())
}

func (i *repositoryManager) PelangganRepository() repository.PelangganRepository {
	return repository.NewPelangganRepository(i.infra.SqlDB())
}

func (i *repositoryManager) NotaRepository() repository.NotaRepository {
	return repository.NewNotaRepository(i.infra.SqlDB())
}

func (i *repositoryManager) MejaRepository() repository.MejaRepository {
	return repository.NewMejaRepository(i.infra.SqlDB())
}

func (i *repositoryManager) RincianRepository() repository.RincianRepository {
	return repository.NewRincianRepository(i.infra.SqlDB())
}

func (i *repositoryManager) AuthRepository() repository.AuthRepository {
	return repository.NewAuthRepository(i.infra.SqlDB())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
