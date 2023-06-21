package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

var db, err = sql.Open("mysql", "root:root@/golang_developer?charset=utf8")

func main() {
	// stmt, err := db.Prepare("Insert into posts(title, body) values(?,?)")
	// checkErr(err)

	// result, err := stmt.Exec("My third post", "My third content")
	// checkErr(err)
	// if result != nil {
	// 	fmt.Println("Post created!")
	// 	fmt.Println(result)
	// }
	rows, err := db.Query("Select * from posts")
	checkErr(err)

	// Slice of items to add all posts
	items := []Post{}

	for rows.Next() {
		post := Post{}

		rows.Scan(&post.Id, &post.Title, &post.Body)
		items = append(items, post)
		fmt.Println(post.Id, post.Title, post.Body)
	}
	db.Close()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("./templates/index.html"))
		if err := t.ExecuteTemplate(rw, "index.html", items); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
