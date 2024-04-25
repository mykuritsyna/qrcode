package greeting

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChooseLanguage() (lang string) {
	fmt.Println("ru - 1, en - 2")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		switch command {
		case "ru":
			lang = "ru"
			return
		case "en":
			lang = "en"
			return
		default:
			fmt.Println("Неверная команда. Пожалуйста, выберите команду из меню." +
				"\n Unknown command. Please choose a command from menu")
		}
	}
	return lang
}
