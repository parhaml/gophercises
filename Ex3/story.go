package cyoa

import (
	"encoding/json"
	"io"
)

// JSONStory - Parse json format of a story to Story type
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	return story, nil
}

// Story - container for entire story.json
type Story map[string]Chapter

// Chapter - struct of the standard format for each chapter of story.json
type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"story,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

// Option - struct of choices for next Chapter in story.json
type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"arc,omitempty"`
}
