package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to mytime")

	presentTime := time.Now()
	fmt.Println("The time is", presentTime.Format("3:04:03 on Monday, January 2, 2006"))

}
