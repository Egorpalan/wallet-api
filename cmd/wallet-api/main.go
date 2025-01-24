package main

import (
	"context"
	"github.com/Egorpalan/wallet/internal/handler"
	"github.com/Egorpalan/wallet/internal/repository"
	"github.com/Egorpalan/wallet/internal/service"
	"github.com/Egorpalan/wallet/pkg/config"
	"github.com/Egorpalan/wallet/pkg/database"
	"github.com/Egorpalan/wallet/pkg/server"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/Egorpalan/wallet/docs"
)

// @title Wallet API
// @version 1.0
// @description This is a wallet management API
// @host localhost:8080
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.LoadConfig("config.env")
	if err != nil {
		logrus.Fatalf("error initializing configs %s", err.Error())
	}

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg)

	srv.Router.POST("/api/v1/wallet", handlers.HandleWalletOperation)
	srv.Router.GET("/api/v1/wallets/:walletId", handlers.GetBalance)
	srv.Router.POST("/api/v1/wallets", handlers.CreateWallet)

	srv.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Fatalf("Failed to start server: %s", err.Error())
		}
	}()
	logrus.Info("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	logrus.Info("App shutdown")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Failed to shutdown server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Failed to close database: %s", err.Error())
	}
}
