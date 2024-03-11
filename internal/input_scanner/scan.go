package input_scanner

import (
	"fmt"
	"strings"
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
