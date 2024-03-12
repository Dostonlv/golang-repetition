package main

import "fmt"

func main() {
	var myString string
	fmt.Println(myString) // myString is an empty string cause it's a zero value
	myString = "Hello Gophers!"
	fmt.Println(myString) // myString is now "Hello Gophers!"

	var myInt int
	fmt.Println(myInt) // myInt is 0 cause it's a zero value
	myInt = 42
	fmt.Println(myInt) // myInt is now 42

	var myBool bool
	fmt.Println(myBool) // myBool is false cause it's a zero value
}
