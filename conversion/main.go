package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("testing", input)

	fmt.Println("welcome to our pizza application")
	fmt.Println("please rate our pizza between 1 and 5: ")
	rating, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", numRating+1)
	}

}
