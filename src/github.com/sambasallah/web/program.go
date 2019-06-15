package main

import(
	"fmt"
)

func main() {
	age := 22
	var num *int

	num = &age

	fmt.Println(*num)


}