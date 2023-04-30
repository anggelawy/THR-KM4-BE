package main
import "fmt"

func removeLeftRight(arr []any, position string) []any {
if position == "left" {
	return arr[1:]

} else if position == "right" {
	return arr[:len(arr)-1]
}
	return arr
}

func main() {
	arr1 := []any{1, 2, 3, 4, 5}
    fmt.Printf("input: %v\noutput: %d\n", arr1, removeLeftRight(arr1,"left"))

	arr2 := []any{1, 2, 3, 4, 5}
    fmt.Printf("input: %v\noutput: %d\n", arr2, removeLeftRight(arr2,"right"))

	arr3 := []any{"\"Edo\",", "\"Budi\",", "\"Joko\",", "\"Tono\""}
    fmt.Printf("input: %v\noutput: %v\n", arr3, removeLeftRight(arr3,"left"))

	arr4 := []any{"\"Edo\",", "\"Budi\",", "\"Joko\",", "\"Tono\""}
    fmt.Printf("input: %v\noutput: %s\n", arr4, removeLeftRight(arr4,"right"))

}
