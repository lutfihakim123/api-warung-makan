package manager

import "api-warung-makan/repository"

type RepositoryManager interface {
	MenuRepository() repository.MenuRepository
	KaryawanRepository() repository.KaryawanRepository
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

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
