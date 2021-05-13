package main
import "fmt"
func main() {
	Calculate(func(a int, b int) int {
		return a + b
	})
}
func Calculate(func1 func(int, int) int) {
	result := func1(10, 20)
	fmt.Println(result)
}
// func add(a int, b int) int {
// 	return a + b
// }