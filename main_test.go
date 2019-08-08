package main

import (
	"gotest.tools/assert"
	"testing"
)

func runstack(t *testing.T, before []string, after []string) {
	env := map[string][]string{}
	runstackWithEnv(t, before, after, env)
}

func runstackWithEnv(t *testing.T, before []string, after []string, env map[string][]string) {
	var stack []string
	stack, _ = consume(stack, before, "", env)
	assert.DeepEqual(t, stack, after)
}

func TestAdd(t *testing.T) {
	runstack(t, []string{"5", "6", "+"}, []string{"11"})
}

func TestSub(t *testing.T) {
	runstack(t, []string{"7", "3", "-"}, []string{"4"})
}

func TestDup(t *testing.T) {
	runstack(t, []string{"5", "dup"}, []string{"5", "5"})
}

func TestDrop(t *testing.T) {
	runstack(t, []string{"5", "6", "drop"}, []string{"5"})
}

func TestSwap(t *testing.T) {
	runstack(t, []string{"5", "6", "swap"}, []string{"6", "5"})
}

func TestOver(t *testing.T) {
	runstack(t, []string{"5", "6", "over"}, []string{"5", "6", "5"})
}

func TestRot(t *testing.T) {
	runstack(t, []string{"5", "6", "7", "rot"}, []string{"6", "7", "5"})
}

func TestMultiple(t *testing.T) {
	runstack(t, []string{"5", "6", "+", "dup", "-"}, []string{"0"})
}

func TestLookup(t *testing.T) {
	env := map[string][]string{
		"plus5": parse("5 +"),
	}
	runstackWithEnv(t, []string{"6", "plus5"}, []string{"11"}, env)
}

func TestLookupTwice(t *testing.T) {
	env := map[string][]string{
		"plus5": parse("5 +"),
	}
	runstackWithEnv(t, []string{"0", "plus5", "plus5"}, []string{"10"}, env)
}

func TestDefineWord(t *testing.T) {
	runstack(t, parse(": foo 1 2 + ; foo"), []string{"3"})
	runstack(t, parse("1 : foo 2 + ; foo"), []string{"3"})
	runstack(t, parse("1 2 : foo + ; foo"), []string{"3"})
}

func TestRedefineWord(t *testing.T) {
	runstack(t, parse(": foo 1 ; foo : foo 2 ; foo"), []string{"1", "2"})
}

func TestDefineIsReentrant(t *testing.T) {
	var env = make(map[string][]string)
	var compile = ""
	var stack []string
	stack, compile = consume(stack, parse("2 : foo"), compile, env)
	assert.DeepEqual(t, stack, []string{"2"})
	assert.Equal(t, compile, "foo")
	stack, compile = consume(stack, parse("1 ; foo"), compile, env)
	assert.Equal(t, compile, "")
	assert.DeepEqual(t, stack, []string{"2", "1"})
	assert.DeepEqual(t, env["foo"], []string{"1"})
}

func TestParse(t *testing.T) {
	assert.DeepEqual(t, parse("5 6 + dup -"),
		[]string{"5", "6", "+", "dup", "-"})
}
