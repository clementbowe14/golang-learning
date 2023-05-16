package main

import "fmt"

type Code string

func(c Code) Int() int8 {
	if c == "something"{
		return 1
	} 

	return 0
}

func main() {




fmt.Println(Code.Int("Something"))

}
