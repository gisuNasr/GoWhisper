package main

import (
	"fmt"
	"net/http"

	"github.com/gisuNasr/GoWhisper/internal/config"
	"github.com/gisuNasr/GoWhisper/internal/infrastructure"

	"log"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := infrastructure.InitDb(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer db.Close()

	err = http.ListenAndServe(cfg.Server.Addr(), nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	fmt.Println("application is up and running")

}
