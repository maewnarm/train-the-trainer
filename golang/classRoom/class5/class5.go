package class5

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

/*type Employee struct {
	Name string
	Dept string
}*/

type EmpList struct {
	Lists []Employee
}

type Employee struct {
	Name       string
	Department []Dept
	Grade      int
}

type Dept struct {
	Name         string
	Section      string
	Section_code string
}

func Mainclass5() {
	/*emp := map[string]string{
		"Name": "",
		"Dept": "",
	}
	content := "<strong> {{.Name}}</strong>\n<strong>{{.Dept}}</strong>\n"

	t := template.New("html")
	t, err := t.Parse(content)
	if err != nil {
		log.Fatalln("parse error !!!", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Failed to render template", err)
	}

	emp["Name"] = "Voramet"
	emp["Dept"] = "PE"
	err = t.Execute(os.Stdout, emp)
	if err != nil {
		log.Fatalln("Failed to render template", err)
	}
	emp["Name"] = "Arraya"
	emp["Dept"] = "QA"
	err = t.Execute(os.Stdout, emp)
	if err != nil {
		log.Fatalln("Failed to render template", err)
	}*/

	/*infile, err := os.Open("./classRoom/class5/index.html")
	defer infile.Close()
	if err != nil {
		log.Fatalln("Read file error", err)
		return
	}

	outfile, err := os.Create("./classRoom/class5/index-1.html")
	defer outfile.Close()
	if err != nil {
		log.Fatalln("Create file error", err)
		return
	}
	_, err = io.Copy(outfile, infile)
	if err != nil {
		log.Fatalln("Copy file error", err)
		return
	}*/

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	/*r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
			log.Println(r.URL.Path)

		})
	}) */

	r.Handle("/emp_list", http.HandlerFunc(handleEmpList))
	r.Handle("/input", http.HandlerFunc(handleInput))

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalln(err)
	}
}

func handleEmpList(w http.ResponseWriter, r *http.Request) {
	data := EmpList{
		Lists: []Employee{
			{
				Name: "Voramet",
				Department: []Dept{
					{Name: "PE", Section: "PE2", Section_code: "4674"},
				},
				Grade: 8,
			},
			{
				Name: "Arraya",
				Department: []Dept{
					{Name: "QA", Section: "QA-New model", Section_code: "4811"},
				},
				Grade: 8,
			},
		},
	}

	content, err := ioutil.ReadFile("./classRoom/class5/index.html")
	if err != nil {
		log.Println("error loading file", err)
		return
	}
	t := template.New("html")
	t, _ = t.Parse(string(content))
	t.Execute(w, data)
}

func handleInput(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Println(q)
	fmt.Println(q["q"])
	fmt.Println(q["q1"])
	io.Copy(w, r.Body)
}
