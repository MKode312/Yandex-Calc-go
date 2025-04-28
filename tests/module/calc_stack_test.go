package module_test

import (
	"testing"

	calc "calculator_go/internal/utils/agent/calculation"
)

func TestCalcStack(t *testing.T) {
	CalcStack := &calc.Stack{}

	// Проверка, что стек пустой
	if len(CalcStack.Items) != 0 {
		t.Errorf("Expected stack to be empty, but it is not.")
	}

	// Тестирование Push
	CalcStack.Push(1.0)
	if len(CalcStack.Items) != 1 {
		t.Errorf("Expected stack size to be 1 after pushing an element.")
	}
	if top := CalcStack.Pop(); top != 1.0 {
		t.Errorf("Expected top of stack to be 1.0, but got %v", top)
	}

	CalcStack.Push(2.5)
	if len(CalcStack.Items) != 1 {
		t.Errorf("Expected stack size to be 1 after pushing another element.")
	}
	if top := CalcStack.Pop(); top != 2.5 {
		t.Errorf("Expected top of stack to be 2.5, but got %v", top)
	}

	// Проверка на пустом стеке
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when popping from empty stack, but did not panic.")
		}
	}()
	CalcStack.Pop() // Пытаемся вытащить из пустого стека
}

func TestCalcStackMultiplePushPop(t *testing.T) {
	CalcStack := &calc.Stack{}

	valuesToPush := []float64{10.0, 20.0, 30.0}
	for _, v := range valuesToPush {
		CalcStack.Push(v)
	}

	for i := len(valuesToPush) - 1; i >= 0; i-- {
		value := CalcStack.Pop()
		if value != valuesToPush[i] {
			t.Errorf("Expected popped value to be %v, but got %v", valuesToPush[i], value)
		}
	}

	if len(CalcStack.Items) != 0 {
		t.Errorf("Expected stack to be empty after popping all elements.")
	}
}