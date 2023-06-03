package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrzalr/printworks/internal/server"
	"github.com/mrzalr/printworks/pkg/db/mysql"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when loading .env files | err = %v", err)
	}

	db, err := mysql.New()
	if err != nil {
		log.Fatalf("error when connecting to database | err = %v", err)
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("error when auto migrating database | err = %v", err)
	}

	s := server.New(db)
	err = s.Run()
	if err != nil {
		log.Fatalf("error when running the server | err = %v", err)
	}
}
