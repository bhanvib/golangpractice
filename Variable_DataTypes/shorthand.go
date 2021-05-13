package main

// import "fmt"
// import "reflect"
import (
	"fmt"
	"reflect"
)

//using shorthand syntax we can declare only local variables
func main() {
	name := "Mukta"
	id := 100
	status, address, code := true, "Pune", 123
	fmt.Println(status, address, code, name, id)
	fmt.Printf("%T \n", status)
	fmt.Printf("%T \n", address)
	fmt.Printf("%T \n", code)
	fmt.Println(reflect.TypeOf(status))
}
