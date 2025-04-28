package calculation

import (
	"fmt"
	"strconv"
	"strings"
)

func Evaluate(expr string) (float64, error) {
	var stack Stack
	tokens := strings.Split(expr, " ")

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			if len(stack.Items) < 2 {
				return 0, fmt.Errorf("not enough operands")
			}
			op1 := stack.Pop()
			op2 := stack.Pop()
			ans, err := Calculate(op1, op2, token)
			if err != nil {
				return 0, err
			}
			stack.Push(ans)
		} else {
			op, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack.Push(op)
		}
	}

	if len(stack.Items) != 1 {
		return 0, fmt.Errorf("too many operands")
	}

	return stack.Pop(), nil
}

// Calculate - вычисляет
func Calculate(op1 float64, op2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return op1 / op2, nil
	default:
		return 0, fmt.Errorf("invalid operator: %s", operator)
	}
}