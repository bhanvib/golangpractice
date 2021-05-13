package main
import "fmt"
func main() {
	// defer foo()
	// bar()
	ConnectDB()
}
func ConnectDB() int {
	defer cleanup()
	status := false
	if status == true {
		return 100
	} else {
		return 200
	}
}
func cleanup() {
	fmt.Println("Do the cleanup")
}
func foo() {
	fmt.Println("foo is called")
}
func bar() {
	fmt.Println("bar is called")
}
