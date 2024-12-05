package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func performOperation(stack *[]string, operation func(a, b int) int) { // Do required operation (+-/*)
	if len(*stack) < 2 {
		log.Println("Error: Not enough elements in the stack for operation")
		return
	}
	s0, err1 := strconv.Atoi((*stack)[len(*stack)-2])
	s1, err2 := strconv.Atoi((*stack)[len(*stack)-1])
	if err1 == nil && err2 == nil {
		*stack = (*stack)[:len(*stack)-2]
		*stack = append(*stack, strconv.Itoa(operation(s0, s1)))
	} else {
		log.Println("Error: Stack elements are not integers")
	}
}

func innerexec(s string) {
	var stack []string // Stack
	var fn string      // Function storage
	var i int          // Current index in the input string

	for i < len(s) {
		if s[i] == ';' {
			break
		} // Stop execution if ";" is encountered

		switch s[i] {
		case '>': // Add next character to stack
			i++
			if i < len(s) {
				stack = append(stack, string(s[i]))
			}
		case '^': // Print stack
			log.Println(strings.Join(stack, ""))
		case '!': // Print element from stack with the index of the following number
			if i+1 < len(s) {
				index, err := strconv.Atoi(string(s[i+1]))
				if err == nil && index < len(stack) {
					log.Println(stack[index])
				} else {
					log.Println("Error: Invalid index for stack access")
				}
				i++
			}
		case '<': // Request and add user input to stack
			var input string
			log.Println(".- input: ")
			fmt.Scanln(&input)
			stack = append(stack, input)
		case '+': // Add first two elements in the stack
			performOperation(&stack, func(a, b int) int { return a + b })
		case '-': // Subtract first two elements in the stack
			performOperation(&stack, func(a, b int) int { return a - b })
		case '/': // Divide first two elements in the stack
			performOperation(&stack, func(a, b int) int {
				if b == 0 {
					log.Println("Error: Division by zero")
					return 0
				}
				return a / b
			})
		case '*': // Multiply first two elements in the stack
			performOperation(&stack, func(a, b int) int { return a * b })
		case '#': // Clear stack
			stack = nil
		case '~': // Run stored function
			innerexec(fn)
		case '[': // Read text until a ']' and append it to the stack as a single element
			var m strings.Builder
			i++
			for i < len(s) && s[i] != ']' {
				m.WriteByte(s[i])
				i++
			}
			stack = append(stack, m.String())
		case '{': // Read new function until a '}' and write it to fn
			i++
			var m strings.Builder
			for i < len(s) && s[i] != '}' {
				m.WriteByte(s[i])
				i++
			}
			fn = m.String()
		case '?': // 'If' logic
			i++
			condition := strings.Builder{}
			// Parse condition until it sees the function bracket '{'
			for i < len(s) && s[i] != '[' {
				condition.WriteByte(s[i])
				i++
			}

			// Split the condition into indices
			parts := strings.Split(condition.String(), "=")
			if len(parts) == 2 {
				index1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				index2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

				if err1 == nil && err2 == nil && index1 < len(stack) && index2 < len(stack) {
					if stack[index1] == stack[index2] {
						// Execute commands inside `{}` if condition is true
						var commands strings.Builder
						i++
						for i < len(s) && s[i] != '}' {
							commands.WriteByte(s[i])
							i++
						}
						innerexec(commands.String())
					}
				} else {
					log.Println("Error: Invalid indices or stack access in condition")
				}
			}
		}
		i++
	}
}

func main() {
	var see string
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read user input
	log.Println(".- Type 'quit' to exit")
	for {
		log.Print(".- zorml 2.3 > ")
		if !scanner.Scan() { // Read  user input
			break
		}
		see = scanner.Text()
		if see == "quit" {
			break
		}
		innerexec(see)
	}
}
