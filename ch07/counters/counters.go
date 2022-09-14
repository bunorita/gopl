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

type LineCounter int

var _ io.Writer = (*LineCounter)(nil)

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return len(p), scanner.Err()
}

// ex7.2
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := newCountingWriter(w)
	return cw, &cw.count
}

var _ io.Writer = (*countingWriter)(nil)

func newCountingWriter(w io.Writer) *countingWriter {
	return &countingWriter{w: w}
}

type countingWriter struct {
	count int64
	w     io.Writer
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.count += int64(n)
	return n, err
}
