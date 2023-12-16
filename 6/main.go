package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
				n := childArray[i].Name
				childArray[i].Age = childArray[j].Age
				childArray[i].Name = childArray[j].Name
				childArray[j].Age = a
				childArray[j].Name = n
			}
		}
	}
	fmt.Println(childArray)

	var childData []Children
	byteData, _ := os.ReadFile("childs.json")
	json.Unmarshal(byteData, &childData)

	childData = childArray

	res, _ := json.Marshal(childData)
	os.WriteFile("childs.json", res, 0)

}
