package main

import "fmt"
import "strconv"

var ch = make(chan string, 1)

func main() {
	fmt.Println("==start==")

	var myfunc = func(i int) {
		for {
			fmt.Println("channel " + strconv.Itoa(i) + " is waiting...\n")
			a := <-ch
			fmt.Println("channel " + strconv.Itoa(i) + " is working...")
			fmt.Println("channel:", a)
		}
	}
	//go myfunc(1)
	//go myfunc(2)
	//go myfunc(3)
	//go myfunc(4)
	//go myfunc(5)
	if true {
		go myfunc(6)
	}
	for t := 1; t <= 5; t++ {
		go myfunc(t)
	}

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
