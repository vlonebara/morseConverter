package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Возвращает HTML из файла index.html
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Обрабатывает загрузку файлов и конвертацию
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "[handlers.UploadHandler] Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при парсинге формы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при получении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при чтении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	inputString := strings.TrimSpace(string(fileBytes))
	if inputString == "" {
		http.Error(w, "[handlers.UploadHandler] Файл пуст", http.StatusInternalServerError)
		return
	}

	convertedString, err := service.Convert(inputString)
	if err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().String()
	timestamp = strings.Map(func(r rune) rune {
		if r == ':' || r == ' ' || r == '.' {
			return '_'
		}
		return r
	}, timestamp)

	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}

	outputFileName := fmt.Sprintf("converted_%s%s", timestamp, ext)

	outputDir := "uploads"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при создании директории: "+err.Error(), http.StatusInternalServerError)
		return
	}

	outputPath := filepath.Join(outputDir, outputFileName)

	if err := os.WriteFile(outputPath, []byte(convertedString), 0644); err != nil {
		http.Error(w, "[handlers.UploadHandler] Ошибка при сохранении файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf(
		"Конвертация выполнена успешно!\n\n"+
			"Исходный текст: %s\n\n"+
			"Результат: %s\n\n"+
			"Файл сохранен: %s",
		inputString,
		convertedString,
		outputPath,
	)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
