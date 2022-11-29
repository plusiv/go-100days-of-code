package main

import "fmt"

func main() {
	// Explicit type conversion
	var x int = 10
	var y float32 = 10.5
	fmt.Println(x + int(y))
	fmt.Println(float32(x) + y)

	// Variables declarations
	var v1 int
	var v2, v3 int = 20, 25
	fmt.Println(v1, v2, v3)
	var v4, v5 = 5, "hello"
	fmt.Println(v4, v5)

	// Usin := to declare variables
	v6 := 100.223
	v7, v8 := -0.05, "world"
	fmt.Println(v6, v7, v8)

	// Const variables

}
