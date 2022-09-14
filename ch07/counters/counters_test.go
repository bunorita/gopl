package counters_test

import (
	"os"
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

func TestLineCounter(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		s    string
		want int
	}{
		{"Hello World", 1},
		{"Hello World\nHello World", 2},
		{"Hello World\nHello World\n", 2},
		{"Hello World\nHello World\n\n", 3},
		{"Hello World! こんにちは　世界", 1},
		{"Hello World!\nこんにちは　世界", 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			t.Parallel()

			var c counters.LineCounter
			if _, err := c.Write([]byte(tt.s)); err != nil {
				t.Fatal(err)
			}
			if got := int(c); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCountingWriter(t *testing.T) {
	t.Parallel()

	ss := []string{
		"Hello World\n",
		"How are you?\n",
		"The Go Programming Language\n",
		"プログラミング言語Go\n",
	}

	w, c := counters.CountingWriter(os.Stdout)

	var total int64 = 0

	for _, s := range ss {
		b := []byte(s)
		if _, err := w.Write(b); err != nil {
			t.Fatal(err)
		}
		total += int64(len(b))

		if *c != total {
			t.Errorf("count is %d, want %d", *c, total)
		}
	}
}
