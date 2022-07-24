package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
	// if there is any other type inside of program that has func called getGreeting associated return string
	// then that type is also honorary kind of like promoted automatically to also being of type bot
	// now that you also an honorary member of type bot
	// now can call func called printGreeting
}

func main() {
	eb := englishBot{} // this also type bot
	sb := spanishBot{} // this also type bot

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "hello"
}

func (sp spanishBot) getGreeting() string {
	return "holla"
}