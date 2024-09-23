package main

import (
	"log/slog"
	"os"

	"github.com/leouraltsev/pastebin-golang/cmd/logger"
	app "github.com/leouraltsev/pastebin-golang/internal/app/pastebinapp"
)

func main() {

	log := logger.New()

	if err := app.Run(log); err != nil {
		slog.Error("error init app", slog.String("err_msg:", err.Error()))
		os.Exit(1)
	}
}
