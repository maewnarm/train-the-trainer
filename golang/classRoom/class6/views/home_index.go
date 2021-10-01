package views

import (
	"html/template"
	"net/http"
	"test-golang/classRoom/class6/models"
)

type HomeIndexView struct {
	NameList []models.Employee
}

func RenderHomeIndex(w http.ResponseWriter, data HomeIndexView) error {
	tmpl, err := template.ParseFiles("./classRoom/class6/views/index.html")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	return tmpl.Execute(w, data)
}
