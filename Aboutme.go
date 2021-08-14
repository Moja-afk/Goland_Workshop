package main
import "fmt"

//-Name
//-Age
//-City
//-State

func main() {
	var myName string
	myName = "Moja"

	var myAge int
	myAge = 22

	var myCity string
	myCity = "Baltimore"

	var myState string
	myState = "Maryland"

	fmt. Printf("Hello my name is %s, I am %d years old! I am from %s city %s. Nice to meet you!", myName, myAge, myCity, myState)

}