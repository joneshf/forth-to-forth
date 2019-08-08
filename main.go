package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// vaguely following: https://www.sifflez.org/lectures/ASE/C3.pdf
func main() {
	var stack []string
	var env map[string][]string
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Starting forth-to-forth\n")
	for scanner.Scan() {
		stack = consume(stack, parse(scanner.Text()), env)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parse(input string) []string {
	return strings.Fields(input)
}

func consume(stack, input []string, env map[string][]string) []string {
	for _, word := range input {
		stack = interpret(word, stack, env)
	}
	return stack
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func interpret(word string, stack []string, env map[string][]string) []string {
	switch word {
	case "+":
		var left, right string
		stack, right = pop(stack)
		stack, left = pop(stack)
		parsedLeft, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}
		parsedRight, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}
		stack = append(stack, strconv.Itoa(parsedLeft+parsedRight))
	case "-":
		var left, right string
		stack, right = pop(stack)
		stack, left = pop(stack)
		parsedLeft, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}
		parsedRight, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}
		stack = append(stack, strconv.Itoa(parsedLeft-parsedRight))

	case "dup":
		var right string
		stack, right = pop(stack)
		stack = append(stack, right, right)

	case "drop":
		stack, _ = pop(stack)

	case "swap":
		var first, second string
		stack, first = pop(stack)
		stack, second = pop(stack)
		stack = append(stack, first, second)

	case "over":
		var first, second string
		stack, first = pop(stack)
		stack, second = pop(stack)
		stack = append(stack, second, first, second)

	case "rot":
		var first, second, third string
		stack, first = pop(stack)
		stack, second = pop(stack)
		stack, third = pop(stack)
		stack = append(stack, second, first, third)

	case ".":
		var first string
		stack, first = pop(stack)
		fmt.Println(first)

	case ".s":
		fmt.Printf("<%d> %s\n", len(stack), stack)

		// *, /, mod, =, <, >
		// KEY (-- c) read stdin
		// EMIT (c --) write stdin
		// WORD (-- address length) (also CREATE)
		// NUMBER (-- n)
		// ! (data address --) write
		// @ (address -- data) read
		// BRANCH OFFSET (--) increment IP
		// 0BRANCHH OFFSET (cond --) increments IP
		// NEXT, CALL, DOCOL, EXIT, LIT?
	default:
		instructions, found := env[word]
		if found {
			// Something isn't quite right here.
			// We seem to be restarting with the original stack each time.

			stack = consume(stack, instructions, env)
		} else {
			stack = append(stack, word)
		}
	}
	return stack
}
