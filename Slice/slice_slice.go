package main
import "fmt"
func main() {
	slice1 := []int {10,20,30,40,50,60}
//[L:H]
	fmt.Println(slice1[1:4])  // 1-3: L to H-1
	fmt.Println(slice1[:4])  // start from 0
	fmt.Println(slice1[1:]) // it will be till end
}
