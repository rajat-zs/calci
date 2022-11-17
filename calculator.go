package main

type myError struct {
	message string
}

func (e myError) Error() string {
	return e.message
}

func Addition(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

// divisions
func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, myError{"Division by zero occurred"}
	}
	return a / b, nil
}

func Modulo(a, b int) (int, error) {
	if b == 0 {
		return 0, myError{"Modulo by zero occurred"}
	}

	if b < 0 {
		return -(a % b), nil
	}
	return a % b, nil
}

func calculator(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return Addition(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Division(a, b)
	case "%":
		return Modulo(a, b)
	default:
		return 0, myError{"Invalid operator"}
	}
}
