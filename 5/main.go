package main

import "fmt"

type Children struct {
	Name string
	Age  int
}

var childArray []Children = []Children{
	{
		Name: "child1",
		Age:  12,
	},
	{
		Name: "child2",
		Age:  11,
	},
	{
		Name: "child3",
		Age:  15,
	},
	{
		Name: "child4",
		Age:  3,
	},
	{
		Name: "child5",
		Age:  6,
	},
	{
		Name: "child6",
		Age:  10,
	},
}

func main() {
	for i := 0; i < len(childArray); i++ {
		for j := i + 1; j < len(childArray); j++ {
			if childArray[i].Age > childArray[j].Age {
				a := childArray[i].Age
				childArray[i].Age = childArray[j].Age
				childArray[j].Age = a
			}
		}
	}
	fmt.Println(childArray)

}
