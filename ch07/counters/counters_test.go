package counters_test

import (
	"testing"

	"github.com/bunorita/gopl/ch07/counters"
)

func TestByteCounter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s    string
		want int
	}{
		{"hello", 5},
		{"hello, Dolly", 12},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			t.Parallel()

			var c counters.ByteCounter
			if _, err := c.Write([]byte(tt.s)); err != nil {
				t.Fatal(err)
			}
			if got := int(c); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestWordCounter(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		s    string
		want int
	}{
		{"Hello World", 2},
		{"Hello My World", 3},
		{"Hello My World ", 3},
		{"Hello World! こんにちは　世界", 4},
		{"Hello World!\nこんにちは　世界", 4},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			t.Parallel()

			var c counters.WordCounter
			if _, err := c.Write([]byte(tt.s)); err != nil {
				t.Fatal(err)
			}
			if got := int(c); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
