package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/dmytrodemianchuk/crud-app/internal/config"
	"github.com/dmytrodemianchuk/crud-app/internal/repository/psql"
	"github.com/dmytrodemianchuk/crud-app/internal/service"
	grpc_client "github.com/dmytrodemianchuk/crud-app/internal/transport/grpc"
	"github.com/dmytrodemianchuk/crud-app/internal/transport/rest"
	"github.com/dmytrodemianchuk/crud-app/pkg/database"
	"github.com/dmytrodemianchuk/crud-app/pkg/hash"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "config"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.Name,
		SSLMode:  cfg.DB.SSLMode,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	hasher := hash.NewSHA1Hasher("salt")

	musicsRepo := psql.NewMusics(db)
	musicsService := service.NewMusics(musicsRepo)

	usersRepo := psql.NewUsers(db)
	tokensRepo := psql.NewTokens(db)

	auditClient, err := grpc_client.NewClient(9000)
	if err != nil {
		log.Fatal(err)
	}

	usersService := service.NewUsers(usersRepo, tokensRepo, auditClient, hasher, []byte("sample secret"))

	handler := rest.NewHandler(musicsService, usersService)

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: handler.InitRouter(),
	}

	log.Info("SERVER STARTED")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
