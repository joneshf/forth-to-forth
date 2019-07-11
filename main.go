package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var stack []string
	stack = append(stack, "5")
	stack = append(stack, "6")
	stack = append(stack, "+")
	output := consume(stack)
	fmt.Printf("stack: %#v, output: %#v\n", stack, output)
}

func consume(stack []string) []string {
	var result []string
	for _, word := range stack {
		switch word {
		case "+":
			var left, right string
			result, left = pop(result)
			result, right = pop(result)
			parsedLeft, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}
			parsedRight, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}
			result = append(result, strconv.Itoa(parsedLeft+parsedRight))

		default:
			result = append(result, word)
		}
		log.Printf("end of loop result: %#v\n", result)
	}
	return result
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}
