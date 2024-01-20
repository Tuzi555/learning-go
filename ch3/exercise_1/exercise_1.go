package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	a := greetings[:2]
	b := greetings[1:4]
	c := greetings[4:]
	fmt.Println(a, b, c)
}
