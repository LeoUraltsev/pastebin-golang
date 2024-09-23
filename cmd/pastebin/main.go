package main

import (
	"log/slog"
	"os"

	app "github.com/leouraltsev/pastebin-golang/internal/app/pastebinapp"
)

func main() {
	if err := app.Run(); err != nil {
		slog.Error("error init app", slog.String("err_msg:", err.Error()))
		os.Exit(1)
	}
}