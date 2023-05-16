package main

import (
	"fmt"
	"strconv"
)

func main() {

	// var colors map[string]string
	var isFun bool
	myMap := map[int]string{}

	for i := 0; i < 100; i++ {
		myMap[i] = "Hello" + strconv.Itoa(i)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			delete(myMap, i)
		}
	}

	fmt.Println(myMap)
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000FF",
		"green": "#342342F",
	}

	printMap(colors)

	fmt.Println(isFun)
	fmt.Println(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

	var swap = func(num1 *int, num2 int) {
		(*num1) = num2
	}

	foo := 12
	bar := 15

	swap(&foo, bar)

	fmt.Println(foo)

}


