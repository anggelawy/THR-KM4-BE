package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	isiData := strings.Split(data, ",")
	age, _ := strconv.Atoi(isiData[1])
	return User{
		Name:    isiData[0],
		Age:     age,
		Address: isiData[2],
	}
}

func main() {
	data := "Edo,20,Jakarta"
	user := ConvertData(data)
	fmt.Printf("input: \"%s\"\n", data)
	fmt.Printf("output: { name: \"%s\", age: %d, address: \"%s\" }\n", user.Name, user.Age, user.Address)

	data1 := "Budi,30,Bandung"
	user1 := ConvertData(data1)
	fmt.Printf("input: \"%s\"\n", data1)
	fmt.Printf("output: { name: \"%s\", age: %d, address: \"%s\" }\n", user1.Name, user1.Age, user1.Address)
}
