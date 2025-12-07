package main

import (
	"calendar/api"
	"fmt"
	"net/http"
)

func main() {
	api.InitRoutes()

	port := "8080"
	portString := ":" + string(port)
	fmt.Println("Сервер запущен на порту " + portString)

	err := http.ListenAndServe(portString, nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
