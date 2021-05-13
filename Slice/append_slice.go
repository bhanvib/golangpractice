package main
import "fmt"
func main() {
	// slice1 := []int{70, 80}
	// GetData(10, 20, 30)
	// GetData(slice1...)
	slice1 := []int {10,20,30}
	slice2 := []int {60,70,80}
	slice2 = append(slice2,slice1...)
	fmt.Println(slice2)
}
func GetData(param1 ...int) {
}