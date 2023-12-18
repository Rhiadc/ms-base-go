package main

import (
	"log"
	"os"

	"github.com/Rhiadc/ms-base-go/app/api"
	"github.com/Rhiadc/ms-base-go/config"
	"github.com/Rhiadc/ms-base-go/infra/logger"
)

func main() {
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	appConfig := config.LoadEnvVars()
	l := logger.NewLogger(*appConfig)
	l.InitLogger()
	l.Logger.Info("Starting API...", "pid", os.Getpid())

	_, err := api.NewServer(api.WithConfig(*appConfig))
	if err != nil {
		log.Fatal(err)
	}

}
