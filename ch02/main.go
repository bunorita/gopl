package main

import (
	"fmt"

	"github.com/bunorita/gopl/ch02/tempconv"
)

func main() {
	c := tempconv.FToC(212.0)
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}
