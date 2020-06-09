package main

import "fmt"

//Hello is a function that returns Hello World string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
