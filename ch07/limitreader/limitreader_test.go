package limitreader_test

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/bunorita/gopl/util"
)

func TestLimitReader(t *testing.T) {
	t.Parallel()

	tests := []*struct {
		str     string
		bufSize int
		limit   int
	}{
		{
			str:     "0123456789",
			bufSize: 0,
			limit:   3,
		},
		{
			str:     "0123456789",
			bufSize: 1,
			limit:   0,
		},
		{
			str:     "0123456789",
			bufSize: 10,
			limit:   3,
		},
		{
			str:     "",
			bufSize: 10,
			limit:   3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%q, bufSize=%d, limit=%d", tt.str, tt.bufSize, tt.limit), func(t *testing.T) {
			t.Parallel()

			r := io.LimitReader(strings.NewReader(tt.str), int64(tt.limit))
			// r := limitreader.LimitReader(strings.NewReader(tt.str), int64(tt.limit))

			target := tt.str // copy target string
			if tt.limit < len(tt.str) {
				target = tt.str[:tt.limit]
			}
			wantBuf := wantBuffer(tt.bufSize, target)
			wantN := util.IntMin(tt.bufSize, len(target))

			buf := make([]byte, tt.bufSize)
			n, err := r.Read(buf)

			if err != nil {
				if len(target) == 0 && errors.Is(err, io.EOF) {
					// OK
				} else {
					t.Fatal(err)
				}
			}

			if got, want := n, wantN; got != want {
				t.Errorf("n got: %d, want: %d\n", got, want)
			}

			if got, want := buf, wantBuf; !reflect.DeepEqual(got, want) {
				t.Errorf("buffer got %v, want %v", got, want)
			}
		})
	}
}

func wantBuffer(size int, target string) []byte {
	buf := make([]byte, size)
	for i, b := range []byte(target) {
		if i > size-1 {
			break
		}
		buf[i] = b
	}
	return buf
}
