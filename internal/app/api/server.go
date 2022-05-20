package api

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	log "github.com/sirupsen/logrus"
	config "wotracker-back/configs"
	_ "wotracker-back/docs"
	"wotracker-back/internal/app/api/controller"
	"wotracker-back/internal/app/api/service"
	db "wotracker-back/internal/pkg"
	logger "wotracker-back/internal/pkg/middleware"
)

type Server struct {
	app            *iris.Application
	miscController *controller.MiscController
}

func (s *Server) RegisterDependencies(config config.Config) {
	log.Debug("initialize db")
	//DB
	dbInstance := db.InitDb(config.DBString)
	//Services
	healthService := service.NewHealthService(dbInstance)
	//Controller
	s.miscController = controller.NewMiscController(healthService)
}

func (s *Server) RegisterRoutes() {
	log.Debug("registering misc routing")
	miscApi := s.app.Party("/misc")
	m := mvc.New(miscApi)
	m.Handle(s.miscController)
}

func (s *Server) StartServer(config config.Config) {
	s.app = iris.New()
	log.Debug("loading middleware...")
	s.app.Use(logger.Logger())

	log.Debug("adding swagger")
	swaggerConfig := &swagger.Config{
		URL: "http://localhost:" + config.PortServer + "/swagger/doc.json", //The url pointing to API definition
	}
	s.app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

	s.RegisterDependencies(config)
	s.RegisterRoutes()

	err := s.app.Listen(":" + config.PortServer)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	} else {
		log.Infof("starting server on port: %s", config.PortServer)
	}
}
