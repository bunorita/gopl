package main

import (
	"fmt"
	"os"

	"github.com/bunorita/gopl/ch05/findlinks"
)

func main() {
	for _, link := range findlinks.Links(os.Stdin) {
		fmt.Println(link)
	}
}
