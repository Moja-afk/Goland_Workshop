package main
import (
	"fmt"
)

func main() {
	myArr := []int{5, 23, 98, 100, 4}
	myArr2 := "Moja"
	fmt.Println(myArr2[0])

	//fmt.Println(myArr)
	//fmt.Println(myArr[len(myArr)-1])
	
	//itterate through the array
	for i := 0; i < len(myArr) - 1; i ++ {
		fmt. Println(myArr[i])
	}

}