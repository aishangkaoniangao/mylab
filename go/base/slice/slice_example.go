package main

import "fmt"

func main() {
	var s1 []string
	s1 = make([]string, 5, 5)
	fmt.Println(s1)

	s1[1] = "node 1"
	s1[2] = "node 2"
	fmt.Println(s1)

	var s2 []string = s1[1:3]
	fmt.Println(s2)

	s2[0] = "node 0"

	fmt.Println(s1, s2)
}
