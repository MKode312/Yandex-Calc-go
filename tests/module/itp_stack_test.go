package module_test

import (
	"testing"
	itp "calculator_go/internal/utils/agent/infix_to_postfix"
)

func TestITPStack(t *testing.T) {
	ITPstack := &itp.Stack{}

	// Проверка, что стек пустой
	if !ITPstack.Empty() {
		t.Errorf("Expected stack to be empty, but it is not.")
	}

	// Тестирование Push
	ITPstack.Push(1)
	if ITPstack.Empty() {
		t.Errorf("Expected stack to not be empty after pushing an element.")
	}
	if top := ITPstack.TopFunc(); top != 1 {
		t.Errorf("Expected top of stack to be 1, but got %v", top)
	}

	ITPstack.Push(2)
	if top := ITPstack.TopFunc(); top != 2 {
		t.Errorf("Expected top of stack to be 2, but got %v", top)
	}

	// Тестирование Pop
	value := ITPstack.Pop()
	if value != 2 {
		t.Errorf("Expected popped value to be 2, but got %v", value)
	}
	if top := ITPstack.TopFunc(); top != 1 {
		t.Errorf("Expected top of stack to be 1 after popping, but got %v", top)
	}

	value = ITPstack.Pop()
	if value != 1 {
		t.Errorf("Expected popped value to be 1, but got %v", value)
	}
	if !ITPstack.Empty() {
		t.Errorf("Expected stack to be empty after popping all elements.")
	}

	value = ITPstack.Pop() // Пытаемся вытащить из пустого стека
	if value != nil {
		t.Errorf("Expected popped value from empty stack to be nil, but got %v", value)
	}
}

func TestITPStackMultiplePushPop(t *testing.T) {
	stack := &itp.Stack{}

	valuesToPush := []interface{}{10, 20, 30}
	for _, v := range valuesToPush {
		stack.Push(v)
	}

	for i := len(valuesToPush) - 1; i >= 0; i-- {
		value := stack.Pop()
		if value != valuesToPush[i] {
			t.Errorf("Expected popped value to be %v, but got %v", valuesToPush[i], value)
		}
	}

	if !stack.Empty() {
		t.Errorf("Expected stack to be empty after popping all elements.")
	}
}