package main

import "fmt"

func main() {
	var myName = "Devin"
	fmt.Println("My name is", myName, myName)

	var name string = "Kathy"
	fmt.Println("name = ", name)

	userName := "Izzy"
	fmt.Println("userName = ", userName)

	var sum int 
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part1 = ", part1, "part2 = ", other)

	part2, other := 2, 0
	fmt.Println("part2 = ", part2, "other = ", other)

	sum = part1 + part2 
	fmt.Println("The sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Go"
	)

	fmt.Println("Lesson name:", lessonName)
	fmt.Println("Lesson type:", lessonType)

	word1, word2, _ := "Hello", "World", "!!!"
	fmt.Println(word1, word2)


}
