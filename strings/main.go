package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello from main.go in the strings directory")
	s := "<<ABC>>"
	fmt.Println(s)
	fmt.Println("Length of string '<<ABC>>':", len(s))
	fmt.Println("Length of string in runes '<<ABC>>':", utf8.RuneCountInString(s))

	c := s[3]
	fmt.Println("c:", c)
	fmt.Printf("c: %c (%T)\n", c, c)

	fmt.Println("Pi:", math.Pi)
	
	// Limit the number of decimals
	fmt.Printf("Pi: %.2f\n", math.Pi)

	a, b := 1, "1"
	msg := fmt.Sprintf("a=%#v, b=%#v", a, b)
	fmt.Println(msg)

	// Strings have no methods (e.g. s.Upper()), see the "strings" built-in package
	x := "jason"
	fmt.Println("x:", x)
	y := strings.ToUpper(x)
	fmt.Println("y:", y)

	// Title method in the strings package handles capitalization
	z := strings.Title(x)
	fmt.Println("z:", z)
}