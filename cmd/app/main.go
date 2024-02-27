package main

import (
	"fmt"
	"strings"
)

func hello() {
	fmt.Println("Привет! Я генератор qr кодов. Введи ссылку, для которой ты хочешь сгенерировать qr-код и я отправлю тебе изображение с кодом.")
	fmt.Println("Пожалуйста, введи ссылку: ")
}

func scan(link *string) {
	fmt.Scan(link)
}

func validate(link string) bool {
	if strings.Contains(link, " ") ||
		strings.HasPrefix(link, "https://") ||
		strings.HasPrefix(link, "http://") ||
		strings.HasPrefix(link, "www.") {
		return true
	}
	return false
}

func main() {
	var link string
	hello()
	scan(&link)
	if validate(link) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
