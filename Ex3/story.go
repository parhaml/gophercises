package cyoa

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
