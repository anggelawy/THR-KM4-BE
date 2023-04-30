package main

import (
	"fmt"
)

func removeUnrelated(dataMap map[string]any, key string) map[string]any {

	delete(dataMap, key)
	

	return dataMap
}


func main() {

	nama := map[string]any{"name": "Edo", "age": 20, "address": "Jakarta"}
	fmt.Println("input: ",nama)
	fmt.Println("output: ", removeUnrelated(nama, "address"))

	gerakan := map[string]any{"run": "lari", "jump": "loncat", "swim": "berenang"}
	fmt.Println("input: ",gerakan)
	fmt.Println("output: ", removeUnrelated(gerakan, "jump"))

	angka := map[string]any{"satu": "ichi", "dua": "ni", "tiga": "san", "empat": "yon",  "lima": "go"}
	fmt.Println("input: ",angka)
	fmt.Println("output: ", removeUnrelated(angka, "tiga"))
}
