package main

import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/get", handler)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("starting server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
