package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	db  *gorm.DB
	app *gin.Engine
}

func New(db *gorm.DB) *server {
	return &server{
		db:  db,
		app: gin.Default(),
	}
}

func (s *server) Run() error {
	s.MapHandlers()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5001"
	}

	log.Printf("server is running on port : %s", port)
	return s.app.Run(fmt.Sprintf(":%s", port))
}
