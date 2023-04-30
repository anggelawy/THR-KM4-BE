package main

import "fmt"

func howManyElements(data []any) int {

	return len(data)
}

func main() {

	slicePertama := []any{1, 2, 3, 4, 5}
	fmt.Printf("input: %v\noutput: %d\n", slicePertama, howManyElements(slicePertama))
	
	sliceKedua := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("input: %v\noutput: %d\n", sliceKedua, howManyElements(sliceKedua))

	sliceKetiga := []any{1, 1, 1, 5, 5, 5}
	fmt.Printf("input: %v\noutput: %d\n", sliceKetiga, howManyElements(sliceKetiga))

	sliceKeempat := []any{"\"Edo\",", "\"Budi\",", "\"Joko\",", "\"Tono\""}
	fmt.Printf("input: %v\noutput: %d\n", sliceKeempat, howManyElements(sliceKeempat))

	sliceKelima := []any{"\"Edo\",", "\"Budi\",", "\"Joko\",", "\"Tono\"", "\"Edo\"", "\"Budi\"", "\"Joko\"", "\"Tono\""}
	fmt.Printf("input: %v\noutput: %d\n", sliceKelima, howManyElements(sliceKelima))

	sliceKeenam := []any{true, false, true, false, true, false}
	fmt.Printf("input: %v\noutput: %d\n", sliceKeenam, howManyElements(sliceKeenam))
}
