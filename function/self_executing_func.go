package main
import "fmt"
func main() {
	//Self Executing function : it will be called at a place where it is defined
	// This func will be executed only once
	// For doing one time initialization of members this is useful
	// Abstraction
	// var exp1 = func() {
	// 	fmt.Println("Self executing func is called")
	// }
	// exp1()
	// exp1()
	// exp1()
	// func() {
	// 	fmt.Println("Self executing func is called")
	// }() //call function here only
	var result = func() int {
		return 1000
	}()
	fmt.Println(result)
}