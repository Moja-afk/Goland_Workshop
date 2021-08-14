package main

//Print every number from 0 to a selected number
//for every number evenly divisibly by 3 print "Fizz"
//For every number evenly divisible by 5 print "Buzz"
//For every number evenly divisible by 3 AND 5 print "FizzBuzz"
//For anything else print the number

import "fmt"

func main() {
	FizzBuzz(100)
}

func FizzBuzz(selectedNum int) {
	for num := 0; num <=selectedNum; num++ {
		if num % 3 == 0 && num % 5 ==0 {
			fmt.Println("FizzBuzz")
		} else if num % 3 == 0 { 
			fmt.Println("Fizz")
		} else if num % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}

	}
}
