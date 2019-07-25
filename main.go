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
	var stack []string;
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stack = append(stack, parse(scanner.Text())...)
		output := consume(stack)
		log.Printf("stack: %#v, output: %#v\n", stack, output)
		stack = output
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parse(input string) []string {
	return strings.Fields(input)
}

func consume(stack []string) []string {
	var result []string
	for _, word := range stack {
		switch word {
		case "+":
			var left, right string
			result, right = pop(result)
			result, left = pop(result)
			parsedLeft, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}
			parsedRight, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}
			result = append(result, strconv.Itoa(parsedLeft+parsedRight))
		case "-":
			var left, right string
			result, right = pop(result)
			result, left = pop(result)
			parsedLeft, err := strconv.Atoi(left)
			if err != nil {
				panic(err)
			}
			parsedRight, err := strconv.Atoi(right)
			if err != nil {
				panic(err)
			}
			result = append(result, strconv.Itoa(parsedLeft-parsedRight))

		case "dup":
			var right string
			result, right = pop(result)
			result = append(result, right, right)

		case "drop":
			result, _ = pop(result)

		case "swap":
			var first, second string
			result, first = pop(result)
			result, second = pop(result)
			result = append(result, first, second)

		case "over":
			var first, second string
			result, first = pop(result)
			result, second = pop(result)
			result = append(result, second, first, second)

		case "rot":
			var first, second, third string
			result, first = pop(result)
			result, second = pop(result)
			result, third = pop(result)
			result = append(result, second, first, third)

		case ".":
			var first string
			result, first = pop(result)
			fmt.Println(first)

		case ".s":
			fmt.Printf("<%d> %s\n",len(result),result)

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
			result = append(result, word)
		}
		//log.Printf("end of loop result: %#v\n", result)
	}
	return result
}

func pop(stack []string) ([]string, string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}
