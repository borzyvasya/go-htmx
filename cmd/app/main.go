package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("qwe")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}

		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		//htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// temp, _ := template.New("t").Parse(htmlStr)
		// temp.Execute(w, nil)

		temp := template.Must(template.ParseFiles("index.html"))
		temp.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
