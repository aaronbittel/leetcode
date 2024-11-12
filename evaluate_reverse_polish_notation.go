package main

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		if isOperator(token) {
			n1, n2 := popTwo(&stack)
			switch token {
			case "+":
				stack = append(stack, n1+n2)
			case "-":
				stack = append(stack, n1-n2)
			case "*":
				stack = append(stack, n1*n2)
			case "/":
				stack = append(stack, n1/n2)
			}
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func isOperator(tok string) bool {
	return tok == "+" || tok == "-" || tok == "*" || tok == "/"
}

func popTwo(stack *[]int) (n1, n2 int) {
	n2 = (*stack)[len(*stack)-1]
	n1 = (*stack)[len(*stack)-2]
	*stack = (*stack)[:len(*stack)-2]
	return
}
