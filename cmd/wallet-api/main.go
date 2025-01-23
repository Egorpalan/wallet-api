package main

import (
	"github.com/Egorpalan/wallet/internal/handler"
	"github.com/Egorpalan/wallet/internal/repository"
	"github.com/Egorpalan/wallet/internal/service"
	"github.com/Egorpalan/wallet/pkg/config"
	"github.com/Egorpalan/wallet/pkg/database"
	"github.com/Egorpalan/wallet/pkg/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg)

	srv.Router.POST("/api/v1/wallet", handlers.HandleWalletOperation)
	srv.Router.GET("/api/v1/wallets/:walletId", handlers.GetBalance)
	srv.Router.POST("/api/v1/wallets", handlers.CreateWallet)

	// Запускаем сервер
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
