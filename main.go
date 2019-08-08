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
	var env = make(map[string][]string);
	var compile = false
	scanner := bufio.NewScanner(os.Stdin)
	log.Printf("Starting forth-to-forth\n")
	for scanner.Scan() {
		stack, compile = consume(stack, parse(scanner.Text()), compile, env)
		if compile {
			fmt.Println("compiled")
		} else {
			fmt.Println("ok")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parse(input string) []string {
	return strings.Fields(input)
}

func consume(stack, input []string, compile bool, env map[string][]string) ([]string, bool) {
	var index = 0
	for(index < len(input)) {
		if !compile {
			stack, compile = interpret(input[index], stack, env)
		} else {
			var definition = input[index]
			var rest = input[index+1:]
			return compiled(definition, rest, stack, env)
		}
		index += 1
	}

	return stack, compile
}

func compiled(definition string, input []string, stack []string, env map[string][]string) ([]string, bool) {
	var index = 0;
	var compile = true
	env[definition] = []string {};
	for(index < len(input)) {
		var word = input[index]
		if(compile) {
			if word == ";" {
				compile = false
			} else {
				env[definition] = append(env[definition], word)
			}
		} else {
			return consume(stack, input[index:], compile, env)
		}
		index += 1
	}
	return stack, compile
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func interpret(word string, stack []string, env map[string][]string) ([]string, bool) {
	var compile = false;
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

  case ":":
		compile = true

	default:
		instructions, found := env[word]
		if found {
			// Something isn't quite right here.
			// We seem to be restarting with the original stack each time.

			stack, compile = consume(stack, instructions, compile, env)
		} else {
			stack = append(stack, word)
		}
	}
	return stack, compile
}
