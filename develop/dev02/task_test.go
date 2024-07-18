package main

import (
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isError  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, "qwe45", false},
		{`qwe\45`, "qwe44444", false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`\`, "", true},
		{`\\`, "\\", false},
	}

	for _, test := range tests {
		result, err := extract(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("Ожидалась ошибка для %q, получен nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Получена неожиданная ошибка от %q, получено %v", test.input, err)
			} else if result != test.expected {
				t.Errorf("Для кейса %q ожидалось %q, получено %q", test.input, test.expected, result)
			}
		}
	}
}
