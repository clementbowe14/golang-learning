package main

import "fmt"
 
func main() {
	channel := make(chan string, 2)
	var word string
	channel <- "First message"
	word = <-channel
	channel <- "Second message"
	fmt.Println(word)
}
