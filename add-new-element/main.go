package main

import "fmt"

func AddElement(data []int, element int, position string) []int {

	if position == "up" {
		return append([]int{element}, data...)
	} else if position == "down" {
		return append(data, element)
	} else {
		return data
	}
}

func main() {

data := []int{1, 2, 3, 4, 5}
result := AddElement(data, 6, "up")
fmt.Println("input: ",result) // Output: [6 1 2 3 4 5]

result = AddElement(data, 6, "down")
fmt.Println("output: ",result) // Output: [1 2 3 4 5 6]

// result = AddElement(data, 6, "middle")
// fmt.Println(result)
}
