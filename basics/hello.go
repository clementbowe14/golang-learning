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
}
