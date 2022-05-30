package main

import (
	"fmt"
	"html/template"
	//"log"
	"net/http"
	"strconv"
)
 
func (app application) home(response http.ResponseWriter, request *http.Request) {
    if request.URL.Path != "/" {
        http.NotFound(response, request)
        return
    }

	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	// Используем функцию template.ParseFiles() для чтения файла шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	template, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = template.Execute(response, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(response, "Internal Server Error", 500)
	}
}
 
func (app application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
 
    fmt.Fprintf(w, "Отображение определенной заметки с ID %d...", id)
}
 
func (app application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Метод не дозволен", http.StatusMethodNotAllowed)
        return
    }
 
    w.Write([]byte("Создание новой заметки..."))
}
