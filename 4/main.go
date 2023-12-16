package main

import "fmt"

func main() {
	var ints []int = []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10}
	max := ints[0]
	min := ints[0]
	for i := 0; i < len(ints); i++ {

		var point *int = &ints[i]

		if *point > max {
			max = ints[i]
		}
		if *point < min {
			min = ints[i]
		}
	}
	fmt.Println(min,max)
}
