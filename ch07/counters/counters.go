package counters

import (
	"io"
)

type ByteCounter int

var _ io.Writer = (*ByteCounter)(nil)

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *ByteCounter) Reset() {
	*c = 0
}
