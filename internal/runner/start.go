package runner

import (
	"github.com/sirupsen/logrus"
	"go_task/internal/adapter/delivery/http"
	"go_task/internal/adapter/delivery/rpc"
	"go_task/internal/app"
	"go_task/internal/app/command"
	"go_task/internal/config"
	"go_task/internal/server"
)

func Start() {
	cfg := newConfig()
	application := newApplication(cfg)
	startServer(cfg, application)
}

func newConfig() *config.Config {
	cfg, err := config.New()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to parse config")
	}

	return cfg
}

func newApplication(cfg *config.Config) app.Application {

	calculator, err := rpc.NewCalculatorUuidClient(cfg)
	if err != nil {
		logrus.Fatalf("Error connection to RPC server err: %v", err)
	}


	commands := app.Commands{
		CalculatingUniqueId: command.NewCalculatingUniqueIdHandler(calculator),
	}

	return app.Application{
		Commands: commands,
	}
}

func startServer(cfg *config.Config, application app.Application) {
	logrus.Info("Starting HTTP server on address :8080")

	httpServer := server.New(cfg, http.NewHandler(application))
	err := httpServer.Run()

	logrus.WithError(err).Fatal("HTTP server stopped")
}
