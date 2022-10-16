package manager

import "api-warung-makan/usecase"

type UseCaseManager interface {
	MenuUseCase() usecase.MenuUseCase
}

type useCaseManager struct {
	repo RepositoryManager
}

func (u *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(u.repo.MenuRepository())
}

func NewUseCaseManager(repo RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
