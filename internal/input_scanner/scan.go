package input_scanner

import (
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func Scan(link *string) {
	fmt.Scan(link)
}

func Validate(link string) bool {
	if !strings.Contains(link, " ") &&
		(strings.HasPrefix(link, "https://") ||
			strings.HasPrefix(link, "http://") ||
			strings.HasPrefix(link, "www.")) {
		return true
	}
	return false
}

func MakeFile(link string) {
	outputDir := "qrcode/output"
	qrPath := outputDir + "/qr_code.png"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		fmt.Println("Ошибка создания каталога:", err)
		return
	}
	err := qrcode.WriteFile(link, qrcode.Medium, 256, qrPath)
	if err != nil {
		fmt.Println("Ошибка генерации QR-кода:", err)
		return
	}
	fmt.Println("QR-код успешно сгенерирован и сохранен")
}
