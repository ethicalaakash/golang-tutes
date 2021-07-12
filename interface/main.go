package main

import "fmt"

type bot interface {
	getGreetings() string
	// this says any type in the program with function called "getGreeting" if you eturn string
	// then you are now an honorary member of type "bot"
	//Now that you're a memeber of type "bot", you can now call this function
	// called "printGreeting(b bot)"
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}
func (englishBot) getGreetings() string {
	return "Hi there"
}

func (spanishBot) getGreetings() string {
	return "Hola!"
}
