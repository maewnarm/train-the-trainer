package views

import (
	"html/template"
	"net/http"
)

func RenderHomeNew(w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles("./classRoom/class6/views/index_create.html")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	return tmpl.Execute(w, nil)
}
