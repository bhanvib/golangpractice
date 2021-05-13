package main

import "fmt"

func main() {
	// for i:=0;i<10;i++{
	// 	fmt.Println("Value of i ",i)
	// }
	// var j = 0
	// for j < 10 {
	// 	j++
	// 	fmt.Println("Value of j ", j)
	// }
	var k = 0
	for {
		k++
		if k > 10 {
			break
		}
		if k%2 == 0 {
			continue
		}
		fmt.Println("Value of k ", k)
	}
}
