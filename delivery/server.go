package delivery

import (
	"api-warung-makan/config"
	"api-warung-makan/delivery/controller"
	"api-warung-makan/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	host := ":8888"
	return &appServer{
		useCaseManager: usecase,
		engine:         ginEngine,
		host:           host,
	}
}

func (a *appServer) initHandlers() {
	controller.NewMenuController(a.engine, a.useCaseManager.MenuUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
