package main
import "fmt"
func main() {
	// var i = 100
	// if i == 100 {
	// 	fmt.Println("Value of i ", i)
	// } else if i == 200 {
	// 	fmt.Println("Value of i ", i)
	// } else {
	// 	fmt.Println("Value of i ", i)
	// }
	// value := GetData()
	// if value == 100{
	// }
	if result := GetData(); result == 100 {
		fmt.Println(result)
	}
}
func GetData() int {
	return 100
}