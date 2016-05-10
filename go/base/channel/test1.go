package main

import "fmt"

//test scanf function

func main() {
	var input string
	fmt.Println("==start==")
	for {
		if _, err := fmt.Scanf("%s", &input); err == nil {
			if input == "exit" || input == "quit" {
				break
			} else {
				fmt.Println("input:", input)
			}
		}
	}
	fmt.Println("==over==")
}
