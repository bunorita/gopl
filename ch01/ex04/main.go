package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, filename := range os.Args[1:] {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s in files: %s\n", n, line, fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !strings.Contains(fileNames[line], f.Name()) {
			fileNames[line] += fmt.Sprintf(" %s", f.Name())
		}
	}
	// ignoring input.Err()
}
