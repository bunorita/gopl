package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/bunorita/gopl/ch02/ex02/unitconv"
	"github.com/bunorita/gopl/ch02/tempconv"
)

func main() {
	for _, arg := range getArgs() {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		showTemperature(v)
		showLength(v)
		showWeight(v)
	}

}

func showTemperature(v float64) {
	f := tempconv.Fahrenheit(v)
	c := tempconv.Celsius(v)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f),
		c, tempconv.CToF(c))
}

func showLength(v float64) {
	m := unitconv.Meter(v)
	f := unitconv.Feet(v)
	fmt.Printf("%s = %s, %s = %s\n",
		m, unitconv.MToF(m),
		f, unitconv.FToM(f))
}

func showWeight(v float64) {
	kg := unitconv.Kg(v)
	p := unitconv.Pound(v)
	fmt.Printf("%s = %s, %s = %s\n",
		kg, unitconv.KgToPound(kg),
		p, unitconv.PoundToKg(p))
}

func getArgs() []string {
	// from args
	if len(os.Args) > 1 {
		return os.Args[1:]
	}

	// from stdin
	var args []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		args = append(args, text)
	}
	return args
}
