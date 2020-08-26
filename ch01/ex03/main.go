package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Concat(os.Args[1:]))
}

func Concat(strings []string) string {
	var s, sep string
	for _, str := range strings {
		s += sep + str
		sep = " "
	}
	return s
}
