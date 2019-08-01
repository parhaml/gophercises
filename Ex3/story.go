package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>

<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>

  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}

    <ul>
    {{range .Options}}
      <li><a href=""/{{.Chapter}}">{{.Text}}</li>
    {{end}}
    </ul>
  </body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

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
