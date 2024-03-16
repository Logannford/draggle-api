package main

// fmt is a package that provides I/O functions
// like printing to the console
import (
	"fmt"

	"rsc.io/quote"
)


func main() {
	fmt.Println(quote.Go())
}