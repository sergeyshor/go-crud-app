package main

import (
	"log"
	"go-crud-app/config"
	"go-crud-app/internal/app"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
