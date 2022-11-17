package main

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	testsAddition := []struct {
		n1     int
		n2     int
		output int
	}{
		{3, 2, 5},    // addition of two numbers
		{-3, 3, 1},   // addition of numbers with one negative and one positive
		{-4, -4, -8}, // take both  negative numbers
	}

	for i := range testsAddition {
		actualOutput := Addition(testsAddition[i].n1, testsAddition[i].n2)
		if reflect.DeepEqual(actualOutput, testsAddition[i].output) {
			t.Log("Correct output")
		} else {
			t.Errorf("input at %v index has expected value as %d and actual output as %d", i, testsAddition[i].output, actualOutput)
		}
	}
}
func TestSubtract(t *testing.T) {
	testsSubtract := []struct {
		num1     int
		num2     int
		response int
	}{
		{2, 1, 1},   // subtraction of two numbers
		{-1, 1, -2}, // subtraction with one negative number
		{-2, -2, 0}, // subtraction of two negative numbers
	}

	for i := range testsSubtract {
		actualOutput := Subtract(testsSubtract[i].num1, testsSubtract[i].num2)
		if reflect.DeepEqual(actualOutput, testsSubtract[i].response) {
			t.Log("Correct output")
		} else {
			t.Errorf("input at %v index has expected value as %d and actual output as %d", i, testsSubtract[i].response, actualOutput)
		}
	}
}

func TestMultiply(t *testing.T) {
	testsMultiply := []struct {
		num1     int
		num2     int
		response int
	}{
		{2, 1, 2},   // simple multiplication
		{-1, 1, -1}, // multiplication with one negative number
		{-2, -2, 4}, // multiplication with two negative numbers
	}

	for i := range testsMultiply {
		actualOutput := Multiply(testsMultiply[i].num1, testsMultiply[i].num2)
		if reflect.DeepEqual(actualOutput, testsMultiply[i].response) {
			t.Log("Correct output")
		} else {
			t.Errorf("input at %v index has expected value as %d and actual output as %d", i, testsMultiply[i].response, actualOutput)
		}
	}
}

func TestDivision(t *testing.T) {
	testsDivision := []struct {
		num1     int
		num2     int
		response int
		err      error
	}{
		{2, 1, 2, nil},   // division of two numbers
		{-2, -2, 1, nil}, // division of two negative integers
		{1, 0, 0, myError{"Division by zero occurred"}}, // through error because not divisible by zero
	}

	for i, testCase := range testsDivision {
		actVal, actError := Division(testCase.num1, testCase.num2)

		if actVal != testCase.response {
			t.Errorf("Test(%v): Expected(%v), Got(%v)", i, testCase.response, actVal)
		}
		if actError != testCase.err {
			t.Errorf("Test(%v): Exp(%v), Got(%v)", i, testCase.err, actError)
		}
	}
}

func TestModulo(t *testing.T) {
	testsModulo := []struct {
		num1     int
		num2     int
		response int
		err      error
	}{
		{123, 10, 3, nil},   //Simple modulo operation
		{123, -10, -3, nil}, // Modulo by negative number
		{12, 0, 0, myError{"Modulo by zero occurred"}}, //Modulo by 0 case

	}

	for i, testcase := range testsModulo {
		actualOut, actErr := Modulo(testsModulo[i].num1, testsModulo[i].num2)

		if actualOut != testcase.response {
			t.Errorf("Test(%v): Exp(%v), Got(%v)", i, testcase.response, actualOut)
		}

		if actErr != testcase.err {
			t.Errorf("Test(%v): Exp(%v), Got(%v)", i, testcase.err, actErr)
		}
	}
}
func TestCalculator(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		op       string
		response int
		err      error
	}{
		{5, 5, "+", 10, nil}, //simple case of addition
		{10, 3, "-", 7, nil}, // simple case of subtraction
		{5, 5, "*", 25, nil}, // simple case of multiplication
		{25, 5, "/", 5, nil}, // simple case of division
		{25, 3, "/", 8, nil}, // improper division case
		{25, 0, "/", 0, myError{"Division by zero occurred"}}, // division by zero error
		{123, 10, "%", 3, nil},                                // simple modulo error
		{123, -10, "%", -3, nil},                              // modulo by negative number
		{12, 0, "%", 0, myError{"Modulo by zero occurred"}},   // Modulo by 0 case
		{5, 2, "#", 0, myError{"Invalid operator"}},           // Invalid operator
	}

	for i, testcase := range tests {
		actualOut, actErr := calculator(tests[i].num1, tests[i].num2, tests[i].op)
		if actualOut != testcase.response {
			t.Errorf("Test(%v): Exp(%v), Got(%v)", i, testcase.response, actualOut)
		}

		if actErr != testcase.err {
			t.Errorf("Test(%v): Exp(%v), Got(%v)", i, testcase.err, actErr)
		}

	}
}
