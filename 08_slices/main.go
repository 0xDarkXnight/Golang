package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices tutorial.")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 865
	highScores[2] = 589
	highScores[3] = 724
	// highScores[4] = 234 // will throw an error

	highScores = append(highScores, 234, 634, 987) // won't throw an error

	fmt.Println(highScores)

	sort.IntsAreSorted(highScores) // false
	sort.Ints(highScores)
	fmt.Println(highScores)
	sort.IntsAreSorted(highScores) // true

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
