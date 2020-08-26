package main

import (
	"strings"
	"testing"
)

func TestConcatWithOneElement(t *testing.T) {
	data := []string{"Hello"}
	result := Concat(data)
	if result != "Hello" {
		t.Errorf("Result is %s, but expected 'Hello'\n", result)
	}
}

func TestConcatWithThreeElement(t *testing.T) {
	data := []string{"Hello", "World", "!"}
	result := Concat(data)
	if result != "Hello World !" {
		t.Errorf("Result is %s, but expected 'Hello World !'\n", result)
	}
}

func BenchmarkConcat(b *testing.B) {
	data := strings.Split("Go is an open source programming language that makes it easy to build simple, reliable, and efficient software", " ")
	for i := 0; i < b.N; i++ {
		Concat(data)
	}
}

func BenchmarkJoin(b *testing.B) {
	data := strings.Split("Go is an open source programming language that makes it easy to build simple, reliable, and efficient software", " ")
	for i := 0; i < b.N; i++ {
		strings.Join(data, " ")
	}
}

