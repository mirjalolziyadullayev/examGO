package main

import "fmt"

func main() {
	son := 2100
	var count int
	for i := 1; i < son; i++ {
		if son%i == 0 {
			count++
		}
	}
	fmt.Print(count)
}
