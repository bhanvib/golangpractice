package main
import (
	"fmt"
	"reflect"
)
//to empty interface we can pass any data type member
func findType(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	//fmt.Println(i.(type))
	switch i.(type) {
	case string:
		fmt.Println(i.(string))
	case int:
		fmt.Println(i.(int))
	}
}
func myFunc(param ...interface{}) {
}
type Product struct {
}
func main() {
//	myFunc(10, "My Name", 190.45)
	findType(100)
	findType("Mukta")
	findType(Product{})
	fmt.Println()
}
