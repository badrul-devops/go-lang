package main

import "fmt"

var Z int = 40

func main() {
	var x int = 20
	fmt.Println(x)
	fmt.Printf("variable type is : %T \n", x)
	y := 30
	fmt.Println(x + y)
	fmt.Println(Z)

	// Multiple variable declaration
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a, b, c)

	// Multiple variable declaration with type
	var (
		d int = 10
		e int = 20
		f int = 30
	)
	fmt.Println(d, e, f)

	// Multiple variable declaration with type

	var (
		g, h, i = 10, 20, 30
	)
	fmt.Println(g, h, i)

	// Multiple variable declaration with type

	var (
		j, k, l int = 10, 20, 30
	)
	fmt.Println(j, k, l)

	// STRING
	var name string = "BADRUL ALAM"
	fmt.Println(name)

	// BOOLEAN
	var isTrue bool = true
	fmt.Println(isTrue)

	// FLOAT
	var floatNumber float32 = 3.14
	fmt.Println(floatNumber)

	// COMPLEX
	var complexNumber complex64 = 1 + 2i
	fmt.Println(complexNumber)

	// BYTE
	var byteNumber byte = 255
	fmt.Println(byteNumber)

	// RUNE
	var runeNumber rune = 255
	fmt.Println(runeNumber)

	// ARRAY
	var arrayNumber [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayNumber)

	// SLICE
	var sliceNumber []int = []int{1, 2, 3, 4, 5}
	fmt.Println(sliceNumber)

	// MAP
	var mapNumber map[string]int = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(mapNumber)

	// STRUCT
	type Person struct {
		name string
		age  int
	}
	var structNumber Person = Person{
		name: "BADRUL ALAM",
		age:  25,
	}
	fmt.Println(structNumber)

}
