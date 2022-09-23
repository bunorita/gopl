package limitreader

import (
	"io"
)

func LimitReader(r io.Reader, n int64) io.Reader {
	return r
}
