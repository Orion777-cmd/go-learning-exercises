package main

import (
	"flag"
	"fmt"
	"strings"
)

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))

	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	pointers()
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func pointers(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n{
		fmt.Println()
	}
}