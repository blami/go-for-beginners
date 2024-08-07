package main

// https://templ.guide/

//go:generate templ generate

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(hello()))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
