package module_test

import (
	"testing"
	itp "calculator_go/internal/utils/agent/infix_to_postfix"
)

func TestToPostfix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2 + 2 * 2", "2 2 2 * +"},                     // Простое выражение
		{"(1 + 2) * 3", "1 2 + 3 *"},                   // Сбалансированные скобки
		{"(1 + (2 * 3))", "1 2 3 * +"},                 // Вложенные скобки
		{"1 + 2 - 3 / (4 * (5 - 6))", "1 2 + 3 4 5 6 - * / -"}, // Сложное выражение с вложенными скобками
		{"10 + (20 - (30 / 5))", "10 20 30 5 / - +"},   // Сложное выражение с несколькими операциями
		{"(3 + (4 * (5 - (6 / (7 + 8)))))", "3 4 5 6 7 8 + / - * +"}, // Глубокие вложенные скобки
		{"9", "9"},                                     // Одно число
	}

	for _, test := range tests {
		result := itp.ToPostfix(test.input)
		if result != test.expected {
			t.Errorf("ToPostfix(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}