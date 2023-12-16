package main

import "fmt"

func main() {


	var get string = "qwertyuiopasdfg"

	if (len(get)%3 == 0) {
		fmt.Println("FIZZ")
	} 
	if (len(get)%5 == 0) {
		fmt.Println("BUZZ")
	}
	if (len(get)%3==0 && len(get)%5==0) {
		fmt.Println("FIZZBUZZ")
	} 
}
