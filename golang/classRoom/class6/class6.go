package class6

import (
	"log"
	"net/http"
	"test-golang/classRoom/class6/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Mainclass6() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//r.Handle("/emp_list", http.HandlerFunc(handleEmpList))
	//r.Handle("/input", http.HandlerFunc(handleInput))
	home := controllers.HomeController{}
	r.Get("/", home.Index)
	r.Get("/new", home.New)
	r.Post("/", home.Create)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln(err)
	}
}
