package main

import (
	"testing"
)

func TestProcessInputString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{"Без чисел и экранирования", "abcd", "abcd", false},
		{"С числами", "a4bc2d5e", "aaaabccddddde", false},
		{"Начинается с цифры", "3abc", "", true},
		{"Начинается с цифры", "45", "", true},
		{"Содержит 2 цифры подряд", "aaa10b", "", true},
		{"Содержит 0 повторений символа", "aaa0b", "aab", false},
		{"Пустая строка", "", "", false},
		{"С экранированием", `d\n5abc`, `d\n\n\n\n\nabc`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessInputString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInputString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("ProcessInputString() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
