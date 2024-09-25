package pastebinapp

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leouraltsev/pastebin-golang/internal/config"
	"github.com/leouraltsev/pastebin-golang/internal/data/test"
	"github.com/leouraltsev/pastebin-golang/internal/http/handlers/paste/get"
	"github.com/leouraltsev/pastebin-golang/internal/http/handlers/paste/save"
)

func Run(log *slog.Logger) error {
	// initial config

	cfg := config.New()
	log.Info("config init ", slog.String("host:", cfg.Host), slog.String("port:", cfg.Port), slog.String("postgres_url:", cfg.PostgresURL))

	// initial http server

	r := chi.NewRouter()
	mock := test.MocksData{}
	r.Get("/{hash}", get.PasteByHash(log, mock))
	r.Get("/", get.LastCreationPaste(log, mock))
	r.Post("/paste", save.New(log, mock))

	srvAddress := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	srv := http.Server{
		Addr:    srvAddress,
		Handler: r,
	}
	log.Info("server starting", slog.String("address", srvAddress))

	if err := srv.ListenAndServe(); err != nil {
		log.Error("shutdown serve", slog.String("error", err.Error()))
	}

	return nil
}
