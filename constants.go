package main

import (
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	s := "unconstant"
	fmt.Println(s)
}
