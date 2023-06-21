package main

import (
	"log"
	"os"

	"github.com/Rhiadc/ms-base-go/app/api"
	"github.com/Rhiadc/ms-base-go/config"
	"github.com/Rhiadc/ms-base-go/domain/book/services"
	"github.com/Rhiadc/ms-base-go/infra/db/gorm"
	"github.com/Rhiadc/ms-base-go/infra/logger"
	"github.com/spf13/cobra"
)

func main() {
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	appConfig := config.LoadEnvVars()
	l := logger.NewLogger(*appConfig)
	l.InitLogger()

	l.Logger.Info("Starting API...", "pid", os.Getpid())

	db, err := gorm.NewGormDB(appConfig.DBConnectionHost)
	if err != nil {
		log.Fatal(err)
	}

	bookRepo := gorm.NewBookRepository(db)
	bookServices := services.NewService(bookRepo)

	_, err = api.NewServer(api.WithService(bookServices), api.WithConfig(*appConfig))
	if err != nil {
		log.Fatal(err)
	}

	migrations := &cobra.Command{
		Use:   "migrations",
		Short: "execute database migrations",
		Run: func(cli *cobra.Command, args []string) {
			gorm.AutoMigrate(db)
		},
	}

	seeds := &cobra.Command{
		Use:   "seeds",
		Short: "inserts test data in database",
		Run: func(cli *cobra.Command, args []string) {
			gorm.AutoSeed(db)
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	if appConfig.ENV == "dev" {
		rootCmd.AddCommand(seeds)
	}
	rootCmd.AddCommand(migrations)
	rootCmd.Execute()

}
