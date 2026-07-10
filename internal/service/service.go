package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Определяет тип входной строки и конвертирует ее
func Convert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("[service.Convert] Входная строка не может быть пустой")
	}

	if isMorseCode(input) {
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("[service.Convert] Не удалось распознать код Морзе")
		}
		return result, nil
	}

	result := morse.ToMorse(input)
	if result == "" {
		return "", errors.New("[service.Convert] Не удалось конвертировать текст в код Морзе")
	}
	return result, nil
}

// Проверяет, является ли строка кодом Морзе
func isMorseCode(input string) bool {
	input = strings.TrimSpace(input)

	if input == "" {
		return false
	}

	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' && char != '/' {
			return false
		}
	}

	hasMorseChar := false
	for _, char := range input {
		if char == '.' || char == '-' {
			hasMorseChar = true
			break
		}
	}

	return hasMorseChar
}

// Предоставляет дополнительные опции для конвертации
// forceType может быть "morse", "text" или "" для автоматического определения
func ConvertWithOptions(input string, forceType string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("[service.ConvertWithOptions] Входная строка не может быть пустой")
	}

	switch strings.ToLower(forceType) {
	case "morse":
		result := morse.ToMorse(input)
		if result == "" {
			return "", errors.New("[service.ConvertWithOptions] Не удалось конвертировать текст в код Морзе")
		}
		return result, nil
	case "text":
		result := morse.ToText(input)
		if result == "" {
			return "", errors.New("[service.ConvertWithOptions] Не удалось распознать код Морзе")
		}
		return result, nil
	default:
		return Convert(input)
	}
}

// Проверяет, является ли строка обычным текстом
func IsText(input string) bool {
	return !isMorseCode(input) && strings.TrimSpace(input) != ""
}

// Проверяет, является ли строка кодом Морзе
func IsMorse(input string) bool {
	return isMorseCode(input)
}

// Выполняет конвертацию и паникует при ошибке
func MustConvert(input string) string {
	result, err := Convert(input)
	if err != nil {
		panic(err)
	}
	return result
}
