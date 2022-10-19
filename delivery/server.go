package delivery

import (
	"api-warung-makan/config"
	"api-warung-makan/delivery/controller"
	"api-warung-makan/delivery/middleware"
	"api-warung-makan/manager"
	"api-warung-makan/usecase"
	"api-warung-makan/utils/authenticator"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	authUseCase    usecase.AuthUseCase
	host           string
	tokenService   authenticator.AccessToken
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}

func (s *appServer) initController() {
	publicRoute := s.engine.Group("/enigma")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)
	controller.NewAppController(publicRoute, s.authUseCase, tokenMdw)
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	host := ":8888"
	tokenService := authenticator.NewAccessToken(config.TokenConfig)
	authUseCase := appServer.authUseCase.NewAuthUseCase(tokenService)
	return &appServer{
		useCaseManager: usecase,
		engine:         ginEngine,
		host:           host,
	}
}

func (a *appServer) initHandlers() {
	controller.NewMenuController(a.engine, a.useCaseManager.MenuUseCase())
	controller.NewKaryawanController(a.engine, a.useCaseManager.KaryawanUseCase())
	controller.NewPelangganController(a.engine, a.useCaseManager.PelangganUseCase())
	controller.NewNotaController(a.engine, a.useCaseManager.NotaUseCase())
	controller.NewMejaController(a.engine, a.useCaseManager.MejaUseCase())
	controller.NewRincianController(a.engine, a.useCaseManager.RincianUseCase())
}
