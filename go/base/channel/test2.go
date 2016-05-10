package main

import "fmt"

var ch = make(chan string, 1)

func main() {
	fmt.Println("==start==")

	var myfunc = func() {
		for {
			fmt.Println("channel is waiting...")
			a := <-ch
			fmt.Println("channel:", a)
		}
	}
	go myfunc()

	var input string

	for {
		fmt.Scanf("%s", &input)
		if input == "exit" || input == "quit" {
			break
		} else {
			ch <- input
		}
	}

	fmt.Println("==over==")
}
