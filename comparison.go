package main

import ( "fmt"; "math/rand"; "time")

func main() {
	rand.Seed(time.Now().UnixNano())
	moneyInAccount := rand.Intn(1000)
	print(moneyInAccount)
	print("\n")
	thingWeWantToBuy := 800

	if moneyInAccount > thingWeWantToBuy {
		fmt.Println("treat yourself, buy what you want man")
	} else if moneyInAccount == thingWeWantToBuy {
		fmt.Println("Think on it. You know we broke")
	} else {
		fmt.Println("Card is going to get declined... don't do it")
	}
}
