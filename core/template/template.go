package template

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = map[string]*template.Template{}

func LoadTemplates() {
	files, _ := filepath.Glob("views/*.html")
	for _, file := range files {
		name := filepath.Base(file)
		templates[name] = template.Must(template.ParseFiles(file))
	}
}

func Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "GoRails: Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
