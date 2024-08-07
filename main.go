package main

// https://templ.guide/

//go:generate templ generate

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/blami/go-for-beginners/blog"
)

var repo blog.Repo

func main() {
	repo = blog.Repo{}

	http.Handle("/", templ.Handler(home()))
	http.Handle("GET /post", templ.Handler(post()))
	http.HandleFunc("POST /post", createPost)
	// TODO: show http.Handle("GET /posts", templ.Handler(posts(repo))) and say why it wouldn't work
	http.HandleFunc("GET /posts", showPosts)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("title"))
	fmt.Println(r.FormValue("body"))
	repo.Add(r.FormValue("title"), r.FormValue("body"))
	fmt.Println(repo.Get())

	http.Redirect(w, r, "/posts", http.StatusSeeOther)
}

func showPosts(w http.ResponseWriter, r *http.Request) {
	posts(repo).Render(r.Context(), w)
}
