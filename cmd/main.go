package main

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/codinomello/webjetos-go/db"
)

//go:embed views/*.html
var views embed.FS

// template html

type Projeto struct {
	Título string
	Alunos string
	ODS    string
	Curso  string
	Ano    string
}

func main() {
	db.Connection()
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Renderiza o template views/index.html
		t, err := template.ParseFS(views, "views/*.html")
		if err != nil {
			panic(fmt.Sprintf("Erro ao carregar templates: %v", err))
		}
		if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, "Erro ao renderizar template: "+err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/post-projeto/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		título := r.PostFormValue("title")
		alunos := r.PostFormValue("students")
		ano := r.PostFormValue("year")
		ods := r.PostFormValue("ods")
		curso := r.PostFormValue("course")

		// TODO: Renderiza o template views/post-projeto.html
		template := template.Must(template.ParseFiles("index.html"))
		template.ExecuteTemplate(w, "film-list-element", Projeto{Título: título, Alunos: alunos, Ano: ano, ODS: ods, Curso: curso})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Servidor rodando na porta %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Erro ao inicializar o servidor: %v\n", err)
	}
}
