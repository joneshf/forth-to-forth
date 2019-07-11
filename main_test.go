package main

import (
	"testing"
)
func TestAdd(t *testing.T) {
	var stack []string
	stack = append(stack, "5")
	stack = append(stack, "6")
	stack = append(stack, "+")
	output := consume(stack)
	if output[0] != "11" {
		t.Errorf("stack: %#v, output: %#v\n", stack, output)
	}
}

func TestSub(t *testing.T) {
	var stack []string
	stack = append(stack, "7")
	stack = append(stack, "3")
	stack = append(stack, "-")
	output := consume(stack)
	if output[0] != "4" {
		t.Errorf("stack: %#v, output: %#v\n", stack, output)
	}
}

func TestDup(t *testing.T) {
	var stack []string
	stack = append(stack, "5")
	stack = append(stack, "6")
	stack = append(stack, "+")
	stack = append(stack, "dup")
	stack = append(stack, "-")
	output := consume(stack)
	if output[0] != "0" {
		t.Errorf("stack: %#v, output: %#v\n", stack, output)
	}
}

