package counters

import (
	"bufio"
	"bytes"
	"io"
)

type ByteCounter int

var _ io.Writer = (*ByteCounter)(nil)

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// ex7.1
type WordCounter int

var _ io.Writer = (*WordCounter)(nil)

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return len(p), scanner.Err()
}
