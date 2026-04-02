package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsTextOrNot(s string) string {
	if strings.HasPrefix(s, ".") || strings.HasPrefix(s, "-") {
		return morse.ToText(s)
	} else {
		return morse.ToMorse(s)
	}
}
