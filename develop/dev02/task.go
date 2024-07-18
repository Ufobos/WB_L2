package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var errIncorrectString = errors.New("некорректная строка")

func extract(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	runes := []rune(s)
	var result []rune
	escaped := false

	for i := 0; i < len(runes); i++ {
		val := runes[i]

		if escaped {
			result = append(result, val)
			escaped = false
			continue
		}

		if val == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(val) {
			if i == 0 {
				return "", errIncorrectString
			}
			prev := runes[i-1]
			count, _ := strconv.Atoi(string(val))
			for j := 0; j < count-1; j++ {
				result = append(result, prev)
			}
		} else {
			result = append(result, val)
		}
	}

	if escaped {
		return "", errIncorrectString
	}

	return string(result), nil
}

func main() {
	tests := []string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`, `\`, `\\`}
	for _, test := range tests {
		fmt.Println("Строка для распаковки: ", test)
		result, err := extract(test)
		if err != nil {
			fmt.Printf("Ошибка: %v\n\n", err)
		} else {
			fmt.Printf("Результат: %s\n\n", result)
		}
	}
}
