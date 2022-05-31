package main

import "fmt"

type bot interface {
	getGreeting() string
}
type FrenchBot struct{}
type EnglishBot struct{}
type SpanishBot struct{}

func main() {

	eb := EnglishBot{}
	sb := SpanishBot{}
	fb := FrenchBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(fb)
}

func (EnglishBot) getGreeting() string {
	return "Hi there"
}

func (FrenchBot) getGreeting() string {
	return "Bonjour"
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
