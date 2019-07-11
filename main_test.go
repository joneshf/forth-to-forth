package main

import (
	"testing"
	"gotest.tools/assert"
)

func runstack(t *testing.T, before []string, after []string) {
	assert.DeepEqual(t, consume(before), after)
}

func TestAdd(t *testing.T) {
	runstack(t, []string {"5", "6", "+"}, []string {"11"})
}

func TestSub(t *testing.T) {
	runstack(t, []string {"7", "3", "-"}, []string {"4"})
}

func TestDup(t *testing.T) {
	runstack(t, []string {"5", "dup"}, []string {"5", "5"})
}

func TestSwap(t *testing.T) {
	runstack(t, []string {"5", "6", "swap"}, []string {"6", "5"})
}

func TestOver(t *testing.T) {
	runstack(t, []string {"5", "6", "over"}, []string {"5", "6", "5"})
}

func TestRot(t *testing.T) {
	runstack(t, []string {"5", "6", "7", "rot"}, []string {"6", "7", "5"})
}

func TestMultiple(t *testing.T) {
	runstack(t, []string {"5", "6", "+", "dup", "-"}, []string {"0"})
}
