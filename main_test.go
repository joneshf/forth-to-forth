package main

import (
	"testing"
	"gotest.tools/assert"
)

func TestAdd(t *testing.T) {
	output := consume([]string {"5", "6", "+"})
	assert.DeepEqual(t, output, []string {"11"})
}

func TestSub(t *testing.T) {
	output := consume([]string {"7", "3", "-"})
	assert.DeepEqual(t, output, []string {"4"})
}

func TestDup(t *testing.T) {
	output := consume([]string {"5", "dup"})
	assert.DeepEqual(t, output, []string {"5", "5"})
}

func TestSwap(t *testing.T) {
	output := consume([]string {"5", "6", "swap"})
	assert.DeepEqual(t, output, []string {"6", "5"})
}

func TestOver(t *testing.T) {
	output := consume([]string {"5", "6", "over"})
	assert.DeepEqual(t, output, []string {"5", "6", "5"})
}

func TestRot(t *testing.T) {
	output := consume([]string {"5", "6", "7", "rot"})
	assert.DeepEqual(t, output, []string {"6", "7", "5"})
}

func TestMultiple(t *testing.T) {
	output := consume([]string {"5", "6", "+", "dup", "-"})
	assert.DeepEqual(t, output, []string {"0"})
}
