package handlers

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func HandlerForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile")
	if err != nil {
		return
	}
	result := make([]byte, 32)
	result, err = io.ReadAll(file)
	result = append(result, ' ')
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(result)
	lettersonly := regexp.MustCompile(`^[\p{Cyrillic}\s-]+$`)
	if lettersonly.MatchString(string(result)) {
		w.Write([]byte(morse.ToMorse(string(result))))
	} else {
		w.Write([]byte(morse.ToText(string(result))))
	}
}
