package main

import (
	"net/http"

	"qrcode/internal/greeting"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/info", infoHandler)

	// Запуск сервера на порту 80
	http.ListenAndServe(":80", nil)

	lang := greeting.ChooseLanguage()
	greeter := greeting.NewGreeter(lang)

	greeter.Hello()

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This programme generates a QR code when you enter a link\n"))
}
