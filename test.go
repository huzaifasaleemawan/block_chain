package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Calculator App!")
	fmt.Println("Please enter an expression (e.g., 2 + 3):")
	expression := readInput()

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}

func readInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func calculate(expression string) (float64, error) {
	operands, operator, err := parseExpression(expression)
	if err != nil {
		return 0, err
	}

	switch operator {
	case "+":
		return add(operands[0], operands[1]), nil
	case "-":
		return subtract(operands[0], operands[1]), nil
	case "*":
		return multiply(operands[0], operands[1]), nil
	case "/":
		return divide(operands[0], operands[1])
	case "^":
		return power(operands[0], operands[1]), nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}

func parseExpression(expression string) ([]float64, string, error) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return nil, "", fmt.Errorf("Invalid expression: %s", expression)
	}

	operand1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return nil, "", fmt.Errorf("Invalid operand: %s", parts[0])
	}

	operand2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return nil, "", fmt.Errorf("Invalid operand: %s", parts[2])
	}

	return []float64{operand1, operand2}, parts[1], nil
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division by zero")
	}
	return a / b, nil
}

func power(a, b float64) float64 {
	return math.Pow(a, b)
}
