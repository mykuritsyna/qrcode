package main

import (
	"fmt"

	"qrcode/internal/greeting"
	"qrcode/internal/input_scanner"
)

func main() {
	//функция выводит предложение выбрать язык - просит на ввод значение
	// значение записывается в lang
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
