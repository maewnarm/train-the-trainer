package controllers

import (
	"log"
	"net/http"
	"strconv"
	"test-golang/classRoom/class6/models"
	"test-golang/classRoom/class6/views"
)

type HomeController struct{}

func (h HomeController) Index(w http.ResponseWriter, r *http.Request) {
	lists, err := models.GetAllEmployee()
	if err != nil {
		log.Println("error loading machines", err)
	}

	err = views.RenderHomeIndex(w, views.HomeIndexView{
		NameList: lists,
	})
	if err != nil {
		log.Println("error rendering views", err)
	}
}

func (h HomeController) New(w http.ResponseWriter, r *http.Request) {
	err := views.RenderHomeNew(w)
	if err != nil {
		log.Println("error rendering views", err)
	}
}

func (h HomeController) Create(w http.ResponseWriter, r *http.Request) {
	grade, err := strconv.Atoi(r.PostFormValue("Grade"))
	if err != nil {
		log.Println("error convert Grade input to int", err)
	}
	err = models.AddNewEmployee(models.Employee{
		Name: r.PostFormValue("Name"),
		Department: models.Dept{
			Name:         r.PostFormValue("Dept_name"),
			Section:      r.PostFormValue("Section"),
			Section_code: r.PostFormValue("Section_code"),
		},
		Grade: grade,
	})
	if err != nil {
		log.Println("error retrieving from data", err)
	}
	http.Redirect(w, r, "/", 302)
}
