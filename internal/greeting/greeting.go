package greeting

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"qrcode/internal/input_scanner"
)

func Hello() {
	fmt.Println("Привет! Я генератор qr кодов. Введи ссылку, для которой ты хочешь сгенерировать qr-код и я отправлю тебе изображение с кодом." +
		"\n Тебе доступны команды: stop - завершение программы, info - вывод информации о программе и формате ссылки, qr - начало создания qr-кода")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		switch command {
		case "stop":
			stopProgram()
			return
		case "info":
			showInfo()
		case "qr":
			fmt.Println("Введите ссылку")
			var link string
			input_scanner.Scan(&link)
			if input_scanner.Validate(link) {
				fmt.Println("Ссылка удачно считана")
				input_scanner.MakeFile(link)
			} else {
				fmt.Println("Ссылка некорректна")
			}
		default:
			fmt.Println("Неверная команда. Пожалуйста, выберите команду из меню.")
		}

		showMenu()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

func showMenu() {
	fmt.Println("Доступные команды:")
	fmt.Println("stop - завершение программы")
	fmt.Println("info - вывод информации о программе и формате ссылки")
	fmt.Println("qr - создание QR-кода")
	fmt.Print("Введите команду: ")
}

func stopProgram() {
	fmt.Println("Программа завершена.")
}

func showInfo() {
	fmt.Println("Программа создает QR-код для введенной пользователем ссылки.")
	fmt.Println("Формат ссылки: http://example.com")
	fmt.Println()
}
