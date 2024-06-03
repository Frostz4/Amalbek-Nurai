package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	w.Header().Del("Cache-Control")
	w.Header().Get("Cache-Control")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Привет из Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))
		return
	}
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
