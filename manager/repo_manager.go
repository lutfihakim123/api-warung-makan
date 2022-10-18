package manager

import "api-warung-makan/repository"

type RepositoryManager interface {
	MenuRepository() repository.MenuRepository
	KaryawanRepository() repository.KaryawanRepository
	PelangganRepository() repository.PelangganRepository
	NotaRepository() repository.NotaRepository
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

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
