package counters_test

import (
	"testing"

	"github.com/bunorita/gopl/ch07/counters"
)

func TestByteCounter(t *testing.T) {
	t.Parallel()

	var c counters.ByteCounter

	tests := []struct {
		p    []byte
		want int
	}{
		{
			p:    []byte("hello"),
			want: 5,
		},
		{
			p:    []byte("hello, Dolly"),
			want: 12,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(string(tt.p), func(t *testing.T) {
			t.Parallel()

			c.Reset()
			if _, err := c.Write(tt.p); err != nil {
				t.Fatal(err)
			}
			if got := int(c); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
