package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Hello World!")
	fmt.Println("DATETIME: ", time.Now())
}
