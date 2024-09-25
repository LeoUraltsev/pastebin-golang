package get

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leouraltsev/pastebin-golang/internal/models"
)

type ResponsePaste struct {
	Title string `json:"title"`
	Data  string `json:"data"`
}

type Response struct {
	Status string       `json:"status"` //OK, ERROR
	Paste  models.Paste `json:"paste"`
}

type PasteGetter interface {
	GetPaste(hash string) (models.Paste, error)
	GetLastPaste(limit int) ([]models.Paste, error)
}

const limit = 10

func PasteByHash(log *slog.Logger, pasteGetter PasteGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.paste.get.New"

		log = log.With(
			slog.String("op", op),
		)

		hash := chi.URLParam(r, "hash")
		log.Info("start getting paste", slog.String("hash", hash))
		paste, err := pasteGetter.GetPaste(hash)
		if err != nil {
			log.Error("error getting paste", slog.String("err", err.Error()))
			return
		}

		res, _ := json.Marshal(ResponsePaste{
			Title: paste.Title,
			Data:  paste.Data,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func LastCreationPaste(log *slog.Logger, pasteGetter PasteGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.paste.get.LastCreationPaste"

		log = log.With(
			slog.String("op", op),
		)

		pastes, err := pasteGetter.GetLastPaste(limit)
		if err != nil {
			log.Error("error getting last paste", slog.String("err", err.Error()))
		}

		var res []ResponsePaste
		for _, paste := range pastes {
			res = append(res, ResponsePaste{
				Title: paste.Title,
				Data:  paste.Data,
			})
		}

		marshal, _ := json.Marshal(res)

		w.WriteHeader(http.StatusOK)
		w.Write(marshal)

	}
}
