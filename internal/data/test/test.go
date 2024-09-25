package test

import (
	"github.com/leouraltsev/pastebin-golang/internal/models"
)

type MocksData struct {
}

func (m MocksData) GetPaste(hash string) (models.Paste, error) {
	return models.Paste{
		Title: "test",
		Data:  "test ts sts",
	}, nil
}

func (m MocksData) GetLastPaste(limit int) ([]models.Paste, error) {
	return []models.Paste{
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
		{
			Title: "test",
			Data:  "test ts sts",
		},
	}, nil
}

func (m MocksData) Save(paste models.Paste) (url string, err error) {
	return "test", nil
}
