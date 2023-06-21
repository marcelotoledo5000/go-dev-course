package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		post := Post{Id: 1, Title: "Why use GoLang?", Body: "What is it and why study GoLang? GoLang, also called Go, is an open source programming language supported by Google. As mentioned, it is easy to learn, has a robust standard library and, above all, has a growing ecosystem of communities and tools. Above all, it is very attractive."}

		if title := r.FormValue("title"); title != "" { // Receive a new Title via querie string
			post.Title = title
		}

		t := template.Must(template.ParseFiles("./templates/index.html"))
		// t.ExecuteTemplate(rw, "index.html", post) // Basic mode without verification.
		if err := t.ExecuteTemplate(rw, "index.html", post); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
