package pastebinapp

import (
	"log/slog"

	"github.com/leouraltsev/pastebin-golang/internal/config"
)

func Run() error {
	// initial config
	cfg := config.New()
	slog.Info("config init ", slog.String("host:", cfg.Host), slog.String("port:", cfg.Port), slog.String("postgres_url:", cfg.PostgresURL))

	return nil
}
