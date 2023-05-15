package main

import "fmt"

func main() {
	c := make(chan int32)

	s := []int32{12,13,12,24,15}


	for i := range(s) {
		go addSomething(c, s)
		s[i] += <- c
	}

	for i := range(s) {
		fmt.Printf("My new final value is %d",  s[i])
	}
}

func addSomething(c, s)