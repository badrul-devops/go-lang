package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays")

	var arrylist [5]string
	arrylist[0] = "zero"
	arrylist[1] = "one"
	arrylist[2] = "two"
	arrylist[3] = "three"
	arrylist[4] = "four"

	fmt.Println("arrylist is ", arrylist)
	fmt.Println("array lenth is ", len(arrylist))

	var fruitlist = [3]string{"apple", "orange", "mango"}
	fmt.Println("fruitlist is ", fruitlist)

}
