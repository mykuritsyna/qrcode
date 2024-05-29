package greeting

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"qrcode/internal/input_scanner"
)

type greeter struct {
	Language string
}

func NewGreeter(lang string) greeter {
	g := greeter{Language: lang}
	return g
}

type texts struct {
	Hello            string
	Link             string
	OkLink           string
	FailLink         string
	Menu             string
	InvalidCommand   string
	AvailableCommand string
	Finish           string
	Info             string
}

func GetTexts() map[string]texts {
	ruTexts := texts{
		Hello:            "Привет! Я генератор qr кодов. Введи ссылку, для которой ты хочешь сгенерировать qr-код и я отправлю тебе изображение с кодом.\nТебе доступны команды: stop - завершение программы, info - вывод информации о программе и формате ссылки, qr - начало создания qr-кода",
		Link:             "Введите ссылку",
		OkLink:           "Ссылка удачно считана",
		FailLink:         "Ссылка некорректна",
		Menu:             "Пожалуйста, выберите команду из меню",
		InvalidCommand:   "Неверная команда",
		AvailableCommand: "Доступные команды: stop - завершение программы, info - вывод информации о программе и формате ссылки, qr - создание QR-кода. Введите команду: ",
		Finish:           "Программа завершена",
		Info:             "Программа создает QR-код для введенной пользователем ссылки. Формат ссылки: http://example.com",
	}

	enTexts := texts{
		Hello:            "Hi! I am a QR code generator. Enter the link for which you want to generate a QR code, and I will send you an image with the code.\nYou have the following commands available: stop - terminate the program, info - display information about the program and link format, qr - start creating a QR code",
		Link:             "Enter the link",
		OkLink:           "Link successfully read",
		FailLink:         "The link is incorrect",
		Menu:             "Please choose a command from the menu",
		InvalidCommand:   "Invalid command",
		AvailableCommand: "Available commands: stop - terminate the program, info - display information about the program and link format, qr - create a QR code. Enter command",
		Finish:           "The program has ended",
		Info:             "The program generates a QR code for the link entered by the user. Link format: http://example.com ",
	}

	return map[string]texts{
		"ru": ruTexts,
		"en": enTexts,
	}

}

func (g greeter) Hello() {

	text := GetTexts()[g.Language]
	fmt.Println(text.Hello)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		switch command {
		case "stop":
			g.stopProgram()
			return
		case "info":
			g.showInfo()
		case "qr":
			fmt.Println(text.Link)
			var link string
			input_scanner.Scan(&link)
			if input_scanner.Validate(link) {
				fmt.Println(text.OkLink)
				input_scanner.MakeFile(link)
			} else {
				fmt.Println(text.FailLink)
			}
		default:
			fmt.Println(text.InvalidCommand)
		}

		g.showMenu()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

func (g greeter) showMenu() {
	text := GetTexts()[g.Language]
	fmt.Println(text.Menu)
}

func (g greeter) stopProgram() {
	text := GetTexts()[g.Language]
	fmt.Println(text.Finish)

}

func (g greeter) showInfo() {
	text := GetTexts()[g.Language]
	fmt.Println(text.Info)
}
