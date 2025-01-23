package server

import (
	"github.com/Egorpalan/wallet/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Router: gin.Default(),
		Config: cfg,
	}
}

func (s *Server) Run() error {
	return s.Router.Run(":" + s.Config.ServerPort)
}
