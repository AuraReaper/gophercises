package internal

import (
	"html/template"

	"github.com/AuraReaper/choose_your_own_adventure/pkg/decoder"
)

type Service struct {
	Data     map[string]decoder.Story
	Template *template.Template
}

func NewService(filePath string) *Service {
	storiesData := decoder.GetStoryData(filePath)
	tmpl := template.Must(template.ParseFiles("layout.html"))

	return &Service{
		Data:     storiesData,
		Template: tmpl,
	}
}
