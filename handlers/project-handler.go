package handlers

import (
	"html/template"
	"net/http"
	"time"
)

type Project struct {
	Title  string
	Group  string
	Year   string
	Course string
	ODS    string
}

func HandleProjects(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	group := r.PostFormValue("students")
	year := r.PostFormValue("year")
	course := r.PostFormValue("course")
	ods := r.PostFormValue("ods")

	// TODO: Renderiza o template views/post-projeto.html
	template := template.Must(template.ParseFiles("index.html"))
	template.ExecuteTemplate(w, "film-list-element", Project{Title: title, Group: group, Year: year, ODS: ods, Course: course})
}
