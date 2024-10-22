package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count int

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		number_str := r.Form.Get("count")
		number, err := strconv.Atoi(number_str)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		count = count + number
	} else if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(count)))
	}
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/count", handler)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("starting server...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
