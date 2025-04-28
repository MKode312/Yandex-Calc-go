package module_test

import (
	"calculator_go/internal/utils/agent/validator"
	"testing"
)

func TestIsValidExpression(t *testing.T) {
	tests := []struct {
		expr     string
		expected bool
	}{
		{"", false},                          // Пустая строка
		{"1 + 2", true},                     // Простое выражение
		{"(1 + 2) * 3", true},               // Сбалансированные скобки
		{"(1 + 2", false},                   // Несбалансированные скобки (открывающая)
		{"1 + 2)", false},                   // Несбалансированные скобки (закрывающая)
		{"(1 + (2 * 3))", true},             // Вложенные сбалансированные скобки
		{"1 + a", false},                    // Неверный символ 'a'
		{"1 + 2 - 3 / 4 * (5 - 6)", true},  // Сложное выражение
		{"1 + 2 - (3 / (4 * (5 - 6)))", true}, // Сложное выражение с вложенными скобками
		{"1 + 2 - (3 / (4 * (5 - 6))", false}, // Несбалансированные скобки
	}

	for _, test := range tests {
		result := validator.IsValidExpression(test.expr)
		if result != test.expected {
			t.Errorf("IsValidExpression(%q) = %v; expected %v", test.expr, result, test.expected)
		}
	}
}