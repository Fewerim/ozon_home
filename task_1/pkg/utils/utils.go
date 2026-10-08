package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// ParseStrToInt64 - парсит строку в тип int64
func ParseStrToInt64(str string) (int64, error) {
	const errorResult = -1

	result, err := strconv.Atoi(str)
	if err != nil {
		return errorResult, fmt.Errorf("failed to parse str '%v'", str)
	}

	return int64(result), nil
}

// ParseStringToArgs - парсит строку в аргументы для приложения
func ParseStringToArgs(s string) ([]string, error) {
	var args []string        // полученные аргументы
	var word strings.Builder // для создания фраз
	inQuotes := false        // флаг для проверки, что фраза в кавычках
	wordStarted := false     // флаг для проверки, что аргумент начался

	for _, ch := range s {
		switch {
		case ch == '\'' || ch == '"':
			inQuotes = !inQuotes
			wordStarted = true
		case unicode.IsSpace(ch) && !inQuotes:
			if wordStarted {
				args = append(args, word.String())
				word.Reset()
				wordStarted = false
			}
		default:
			word.WriteRune(ch)
			wordStarted = true
		}
	}

	if inQuotes {
		return nil, fmt.Errorf("not closed quote")
	}
	if wordStarted {
		args = append(args, word.String())
	}

	return args, nil
}
