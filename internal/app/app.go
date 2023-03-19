package app

import (
	"database/sql"
	"go-crud-app/config"
	"log"
	"go-crud-app/pkg/pghelper"

	"github.com/gin-gonic/gin"

	"go-crud-app/internal/infrastructure/repository/postgres"
	v1 "go-crud-app/internal/transport/http/v1"
	"go-crud-app/internal/usecase"
)

func Run(cfg *config.Config) {
	// Repository
	connURL := pghelper.GetConnURL(cfg.PG)

	conn, err := sql.Open("postgres", connURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Use cases
	userUseCase := usecase.NewUserUseCase(
		postgres.NewUserRepo(conn),
	)

	tweetUseCase := usecase.NewTweetUseCase(
		postgres.NewTweetRepo(conn),
	)

	// HTTP server
	r := gin.Default()
	h := v1.NewHandlers(userUseCase, tweetUseCase)
	h.NewRouter(r)
	r.Run()
}
