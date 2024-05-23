package main

import (
	"fmt"
	"net/http"

	"qrcode/internal/greeting"
	"qrcode/internal/input_scanner"
)

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/info", infoHandler)

	// Запуск сервера на порту 8080
	http.ListenAndServe(":80", nil)

	lang := greeting.ChooseLanguage()
	greeter := greeting.NewGreeter(lang)

	var link string

	greeter.Hello()

	if input_scanner.Validate(link) {
		fmt.Println("Ссылка удачно считана")
	} else {
		fmt.Println("Ссылка некорректна")
	}
	if input_scanner.Validate(link) {
		input_scanner.MakeFile(link)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This programme generates a QR code when you enter a link"))
}
