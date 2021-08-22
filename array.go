package main

import "fmt"

func main() {
	var a [5]int
	for i := 0; i <= 5; i++ {
		fmt.Println("a["+fmt.Sprint(i)+"] = ", i)
		a[i] = i

	}
}
