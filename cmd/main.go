package main

import (
	"fmt"
	"github.com/gisuNasr/GoWhisper/internal/config"
	"github.com/gisuNasr/GoWhisper/internal/storage"

	"log"
)

func main() {
	cfg, err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := storage.InitDb(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}

	defer db.Close()

	fmt.Println("application is up and running")
}
