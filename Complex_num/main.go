package main

import (
	"fmt"
)

func main() {
	x := complex(3.4, 2.1)
	y := complex(13.9, 2)

	fmt.Println(x + y)   // addition
	fmt.Println(x - y)   // subtraction
	fmt.Println(x * y)   // multiplication
	fmt.Println(x / y)   // division
	fmt.Println(real(x)) // real part of x
	fmt.Println(imag(x)) // imaginary part of x
}
