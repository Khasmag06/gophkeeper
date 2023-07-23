package main

import (
	"context"
	"github.com/Khasmag06/gophkeeper/config"
	"github.com/Khasmag06/gophkeeper/internal/api"
	recordRepository "github.com/Khasmag06/gophkeeper/internal/repository/record/postgres"
	authRepository "github.com/Khasmag06/gophkeeper/internal/repository/user/postgres"
	"github.com/Khasmag06/gophkeeper/internal/service/auth"
	"github.com/Khasmag06/gophkeeper/internal/service/record"
	decoder2 "github.com/Khasmag06/gophkeeper/pkg/decoder"
	hasher2 "github.com/Khasmag06/gophkeeper/pkg/hasher"
	jwt2 "github.com/Khasmag06/gophkeeper/pkg/jwt"
	"github.com/Khasmag06/gophkeeper/pkg/logger"
	"github.com/Khasmag06/gophkeeper/pkg/postgres"

	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	l, err := logger.New(cfg.Logger.LogFilePath, cfg.Logger.Level)
	if err != nil {
		log.Fatalf("failed to build logger: %s", err)
	}
	defer func() { _ = l.Sync() }()

	ctx := context.Background()

	db, err := postgres.NewDB(ctx, cfg.PG)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	jwt, err := jwt2.New(cfg.JWT.SignKey)
	if err != nil {
		log.Fatal(err)
	}

	decoder, err := decoder2.New(cfg.Decoder.SecretKey)
	if err != nil {
		log.Fatal(err)
	}
	hasher := hasher2.New(cfg.Hasher.Salt)

	authRepo := authRepository.New(db.Pool)
	authService := auth.New(authRepo, hasher, jwt, decoder)
	recordRepo := recordRepository.New(db.Pool)
	recordService := record.New(recordRepo, decoder)

	r := api.NewHandler(authService, recordService, decoder, l)

	l.Fatal(r.Run(cfg.Server.Port))
}
