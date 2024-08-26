package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Makes exact matching in page root
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippets(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("These are your snippets"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's make a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippets)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :42069")
	err := http.ListenAndServe(":42069", mux)
	log.Fatal(err)
}
