package main
import "fmt"
func main() {
	// switch name {
	// case "John":
	// 	fmt.Println("John/ken is executed")
	// 	fallthrough
	// case "Ken":
	// 	fmt.Println("John/ken is executed")
	// case "Ram":
	// 	fmt.Println("Ram is executed")
	// default:
	// 	fmt.Println("default executed")
	// }
	// switch name {
	// case "John","Ken":
	// 	fmt.Println("John/ken is executed")
	// case "Ram":
	// 	fmt.Println("Ram is executed")
	// default:
	// 	fmt.Println("default executed")
	// }
	var name = "Persistent"
	switch {
	case len(name) == 10:
		fmt.Println("Length is 10")
	case len(name) == 11:
		fmt.Println("Length is 11")
	case len(name) == 12:
		fmt.Println("Length is 12")
	}
}