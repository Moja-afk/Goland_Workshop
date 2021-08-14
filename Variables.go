package main

import "fmt"

func main() {
	//variable name
	var myName string
	myName = "Moja"

	fmt.Println(myName)

	myName = "New Name"
	fmt.Println(myName)

	// VAriable num 2
	//when you declare an integer it takes on the value of 0
	var ourNum int = 6
	fmt.Println(ourNum)

	//VAriable num 3
	ourbool := true
	fmt.Println(ourbool)
	fmt.Printf("%T", ourbool)

	//%T print out the data type rather than the value. so if a = 5 %t, it will print integer
	bottles := 5
	//v is for variable. %d int and %s string
	fmt. Printf("\nThere are %v bottles of beer on the wall, and my name is %v", bottles, myName)



}