package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(left, right int, operator string) (int, error) {
	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func main() {
	fmt.Println("Hi Abhishek.Veeramalla, I am a calculator app ....")

	for {
		// Read input from the user
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter any calculation (Example: 1 + 2 (or) 2 * 5 -> Please maintain spaces as shown in example): ")
		text, _ := reader.ReadString('\n')

		// Trim the newline character from the input
		text = strings.TrimSpace(text)

		// Check if the user entered "exit" to quit the program
		if text == "exit" {
			break
		}

		// Split the input into two parts: the left operand and the right operand
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Convert the operands to integers
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		// Perform the calculation
		result, err := calculate(left, right, parts[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Print the result
		fmt.Printf("Result: %d\n", result)
	}
}

