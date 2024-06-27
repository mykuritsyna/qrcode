package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"qrcode/internal/greeting"
	"qrcode/internal/input_scanner"

	"github.com/skip2/go-qrcode"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/info", info)
	http.HandleFunc("/qr", qr)

	// Запуск сервера на порту 80
	go http.ListenAndServe(":80", nil)

	lang := greeting.ChooseLanguage()
	greeter := greeting.NewGreeter(lang)

	greeter.Hello()

}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username parameter is missing", http.StatusBadRequest)
		return
	}
	response := fmt.Sprintf("Hello, %s\n", username)
	w.Write([]byte(response))
}

func info(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This programme generates a QR code when you enter a link\n"))
}

func qr(w http.ResponseWriter, r *http.Request) {

	type QR struct {
		Link string `json:"link"`
	}

	var req QR

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	isValid := input_scanner.Validate(req.Link)
	if !isValid {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерация QR-кода
	qr, err := qrcode.New(req.Link, qrcode.Medium)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	// Преобразование QR-кода в PNG
	png, err := qr.PNG(256)
	if err != nil {
		http.Error(w, "Failed to create PNG", http.StatusInternalServerError)
		return
	}

	// Установка заголовков и отправка PNG-изображения
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(png)
}
