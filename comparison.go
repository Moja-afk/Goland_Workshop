package main

import ( "fmt")
//; "math/rand"

func main() {
	moneyInAccount := 210
	thingWeWantToBuy := 800

	if moneyInAccount > thingWeWantToBuy {
		fmt.Println("treat yourself, buy what you want man")
	} else if moneyInAccount == thingWeWantToBuy {
		fmt.Println("Think on it. You know we broke")
	} else {
		fmt.Println("Card is going to get declined... don't do it")
	}
}
