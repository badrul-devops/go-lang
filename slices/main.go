package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitlist = []string{"apple", "orange", "mango"}
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 867

	highscore = append(highscore, 555, 666, 777)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

}
