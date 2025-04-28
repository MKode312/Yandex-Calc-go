package module_test

import (
	"calculator_go/internal/utils/agent/calculation"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		expr      string
		expected  float64
		expectErr bool
	}{
		{"3 4 +", 7, false},
		{"10 2 -", 8, false},
		{"5 6 *", 30, false},
		{"8 2 /", 4, false},
		{"3 4 + 2 *", 14, false}, // (3 + 4) * 2
		{"10 2 - 3 +", 11, false}, // (10 - 2) + 3
		{"5 1 2 + 4 * +", 17, false}, // (5 + (1 + 2) * 4) = (5 + (3 * 4)) = (5 +12) =17
		{"7 8 + 3 -", 12, false}, // (7 +8) -3
        {"5 /",0,true}, // недостаточно операндов
        {"5 a +",0,true}, // некорректный токен
        {"1 +",0,true}, // недостаточно операндов
        {"1",1,false}, // только один операнд
        {"1 + +",0,true}, // слишком много операторов
    }

	for _, test := range tests {
        result, err := calculation.Evaluate(test.expr)
		
        if test.expectErr {
            if err == nil {
                t.Errorf("Expected error for expression %q but got none.", test.expr)
            }
            continue
        }

        if err != nil {
            t.Errorf("Unexpected error for expression %q: %v", test.expr, err)
            continue
        }

        if result != test.expected {
            t.Errorf("For expression %q: expected %v but got %v.", test.expr, test.expected, result)
        }
    }
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		op1       float64
		op2       float64
		operator  string
		expected  float64
        expectErr bool
    }{
        {3, 4, "+", 7, false},
        {10, 5, "-", 5, false},
        {6, 7, "*", 42, false},
        {8, 2, "/", 4, false},
        {8, 0, "/", 0, true}, // деление на ноль
        {1, -1, "+", 0, false},
        {1.5, -0.5, "*", -0.75, false},
    }

	for _, test := range tests {
        result, err := calculation.Calculate(test.op1, test.op2, test.operator)

        if test.expectErr {
            if err == nil {
                t.Errorf("Expected error for operation %v %s %v but got none.", test.op1, test.operator, test.op2)
            }
            continue
        }

        if err != nil {
            t.Errorf("Unexpected error for operation %v %s %v: %v", test.op1, test.operator, test.op2, err)
            continue
        }

        if result != test.expected {
            t.Errorf("For operation %v %s %v: expected %v but got %v.", test.op1,
                test.operator,
                test.op2,
                test.expected,
                result)
        }
    }
}