package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = fmt.Sprint(3)
	fmt.Println("Slice s :", s)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])

	// fmt.Println("len:", len(s))

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd:", s)

}
