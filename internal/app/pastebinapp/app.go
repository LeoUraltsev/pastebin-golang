package pastebinapp

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leouraltsev/pastebin-golang/internal/config"
)

func Run(log *slog.Logger) error {
	// initial config

	cfg := config.New()
	log.Info("config init ", slog.String("host:", cfg.Host), slog.String("port:", cfg.Port), slog.String("postgres_url:", cfg.PostgresURL))

	// initial http server

	r := chi.NewRouter()

	

	srvAddress := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	srv := http.Server{
		Addr: srvAddress,
		Handler: r,
	}
	log.Info("server starting", slog.String("address", srvAddress))

	if err := srv.ListenAndServe(); err != nil {
		log.Error("shutdown serve", slog.String("error", err.Error()))
	}

	return nil
}
