package main

import (
	"bufio"
	"fmt"
	"os"
)

// 入力が2回以上現れた行の数とその行のテキストを表示する。
// 標準出力から読み込む。
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
