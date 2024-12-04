package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func performOperation(stack *[]string, operation func(a, b int) int) {
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
	var stack []string // Use a slice for stack management
	var fn string      // Function storage
	var i int          // Current index in the input string

	for i < len(s) {
		// Stop execution if ";" is encountered
		if s[i] == ';' {
			break
		}

		switch s[i] {
		case '>':
			i++
			if i < len(s) {
				stack = append(stack, string(s[i]))
			}
		case '^':
			log.Println(strings.Join(stack, ""))
		case '!':
			if i+1 < len(s) {
				index, err := strconv.Atoi(string(s[i+1]))
				if err == nil && index < len(stack) {
					log.Println(stack[index])
				} else {
					log.Println("Error: Invalid index for stack access")
				}
				i++
			}
		case '<':
			var input string
			log.Println(".- input: ")
			fmt.Scanln(&input)
			stack = append(stack, input)
		case '+':
			performOperation(&stack, func(a, b int) int { return a + b })
		case '-':
			performOperation(&stack, func(a, b int) int { return a - b })
		case '/':
			performOperation(&stack, func(a, b int) int {
				if b == 0 {
					log.Println("Error: Division by zero")
					return 0
				}
				return a / b
			})
		case '*':
			performOperation(&stack, func(a, b int) int { return a * b })
		case '#':
			stack = nil // Clear the stack
		case '~':
			innerexec(fn) // Execute stored function
		case '[':
			var m strings.Builder
			i++
			for i < len(s) && s[i] != ']' {
				m.WriteByte(s[i])
				i++
			}
			stack = append(stack, m.String())
		case '{':
			i++
			var m strings.Builder
			for i < len(s) && s[i] != '}' {
				m.WriteByte(s[i])
				i++
			}
			fn = m.String()
		case '?': // Add "if" logic
			i++
			condition := strings.Builder{}
			// Parse condition until '['
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
						// Execute commands inside `[]` if condition is true
						var commands strings.Builder
						i++
						for i < len(s) && s[i] != ']' {
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
	var see = ""
	see = ">1>2?0=1[>3^];" // Example testing
	for see != "quit" {
		log.Print(".- zorml 2.1 >")
		innerexec(see)
		see = "quit"
	}
}
