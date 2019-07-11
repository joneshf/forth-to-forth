package main

import (
	"testing"
	"gotest.tools/assert"
)

func TestAdd(t *testing.T) {
	var stack []string
	stack = append(stack, "5")
	stack = append(stack, "6")
	stack = append(stack, "+")
	output := consume(stack)
	assert.DeepEqual(t, output, []string {"11"})
}

func TestSub(t *testing.T) {
	var stack []string
	stack = append(stack, "7")
	stack = append(stack, "3")
	stack = append(stack, "-")
	output := consume(stack)
	assert.DeepEqual(t, output, []string {"4"})
}

func TestDup(t *testing.T) {
	var stack []string
	stack = append(stack, "5")
	stack = append(stack, "6")
	stack = append(stack, "+")
	stack = append(stack, "dup")
	stack = append(stack, "-")
	output := consume(stack)
	assert.DeepEqual(t, output, []string {"0"})
}

