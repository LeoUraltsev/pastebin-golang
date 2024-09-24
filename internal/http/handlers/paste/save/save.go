package save

import (
	"encoding/json"
	"hash"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/leouraltsev/pastebin-golang/internal/models"
)

type Request struct {
	Url            string `json:"url"`
	Title          string `json:"title"`
	Data           string `json:"data"`
	Status         string `json:"status"`
	ExpirationTime uint   `json:"expiration_time"`
}

type Response struct {
	Status string `json:"status"` //OK, ERROR
	URL    string `json:"url"`
}

type PasteSaver interface {
	Save(paste models.Paste) (url string, err error)
}

func New(log *slog.Logger, pasteSaver PasteSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.paste.save.New"

		w.Header().Set("content-type", "application/json")

		log = log.With(
			slog.String("op", op),
		)

		var req Request

		data, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("failed get request", slog.String("err", err.Error()))
			return
		}

		err = json.Unmarshal(data, &req)
		if err != nil {
			log.Error("failed to decode request", slog.String("err", err.Error()))
			return
		}

		log.Info("request success decode", slog.Any("req", req))

		var paste models.Paste
		if req.Url != "" {
			paste = models.Paste{
				Url:            req.Url,
				Title:          req.Title,
				Data:           req.Data,
				Status:         req.Status,
				ExpirationTime: time.Duration(req.ExpirationTime),
			}
		} else {
			paste = models.Paste{
				//TODO:creation method to generated hash
				Url:            "test",
				Title:          req.Title,
				Data:           req.Data,
				Status:         req.Status,
				ExpirationTime: time.Duration(req.ExpirationTime),
			}
		}

		url, err := pasteSaver.Save(paste)
		if err != nil {
			log.Error("failed to save paste", slog.String("err", err.Error()))
			return
		}

		resp := Response{
			Status: "OK",
			URL:    url,
		}
		res, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusCreated)
		w.Write(res)

	}
}
