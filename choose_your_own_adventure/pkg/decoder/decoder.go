package decoder

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Story struct {
	Key     string   `json:"key"`
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"` // arc -> key
	} `json:"options"`
}

func GetStoryData(filePath string) map[string]Story {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("unable to open file")
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error reading the file, %v", err)
	}

	var stories map[string]Story
	err = json.Unmarshal(byteValue, &stories)
	if err != nil {
		log.Fatalf("error parsing the JSON data, %v", err)
	}

	return stories
}
