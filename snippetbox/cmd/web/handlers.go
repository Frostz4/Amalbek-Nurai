package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errorLog.Println("Некорректный путь:", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	// Логирование текущей рабочей директории
	cwd, err := os.Getwd()
	if err != nil {
		app.errorLog.Println("Ошибка получения текущей директории:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	app.infoLog.Println("Текущая рабочая директория:", cwd)

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	app.infoLog.Println("Попытка парсинга файлов шаблонов:", files)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println("Ошибка парсинга шаблонов:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Попытка выполнения шаблона")
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.errorLog.Println("Некорректный ID:", r.URL.Query().Get("id"))
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.errorLog.Println("Некорректный метод:", r.Method)
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Создание новой заметки..."))
}
