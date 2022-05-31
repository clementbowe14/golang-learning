package main

import (
	"fmt"
)

// type shapable {
// 	model Model
// }

//we are able to create variables outside of functions we just can't assign a value to it
var counter int

func main() {
	fmt.Println("Hi there")
	counter = count(13)
	fmt.Println("count: ", counter)

	//by doing this golang can also infer the type of the variable declared
	fruit := "Pineapple"
	favoriteLanguage := "java"
	worstLanguage := "python"
	isPythonTheWorstLanguage := true

	fmt.Println(fruit, favoriteLanguage, worstLanguage, "and the survey says", isPythonTheWorstLanguage)
}

func count(current int) int {

	return current + 1
}

func sayMyName(name string) {
	fmt.Println(name)
}

func whereIsWaldo() {

	myArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range myArr {

		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}

}
