package main
import "fmt"
type Employee struct {
	FirstName string
	LastName  string
	id        int
}
func main() {
	//Way 1
	emp1 := Employee{FirstName: "Sachin", LastName: "T", id: 100}
	fmt.Println(emp1.FirstName, emp1.LastName)
	//Way 2
	emp2 := Employee{"Virat", "K", 200}
	fmt.Println(emp2.FirstName, emp2.LastName)
	//Way 3 //default value 
	emp3 := Employee{}
	fmt.Println(emp3)
	emp3.FirstName="fknkve"
	fmt.Println(emp3)
}