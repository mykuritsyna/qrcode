package main

import (
	"fmt"

	"qrcode/internal/greeting"
	"qrcode/internal/input_scanner"
)

func main() {
	var link string
	greeting.Hello()
	input_scanner.Scan(&link)
	if input_scanner.Validate(link) {
		fmt.Println("Ссылка удачно считана")
	} else {
		fmt.Println("Ссылка некорректна")
	}
}
