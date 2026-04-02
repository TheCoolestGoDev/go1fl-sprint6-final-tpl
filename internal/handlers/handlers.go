package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Ошибка отображения файла")
		return
	}
	result, err := io.ReadAll(file)
	result = append(result, ' ')
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	resultFile, err := os.Create("resultFile.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла")
		return
	}
	defer resultFile.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	resultFile.WriteString(string(result))
	w.Write(result)
	resultFile.WriteString(service.IsTextOrNot(string(result)))
	w.Write([]byte(service.IsTextOrNot(string(result))))
}
