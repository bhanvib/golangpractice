package main
import "fmt"
func main() {
	Func1("Mukta", "Pune", 123)
	result := Func2("Mukta")
	fmt.Println(result)
	name, _ := GetEMpDetails(10)
	fmt.Println(name)
}
func GetEMpDetails(id int) (string, int) {
	var Name = "Sachin"
	var Salary = 19000
	return Name, Salary
}
//func Func1(Name string,Address string, Code int) {
func Func1(Name, Address string, Code int) {
	fmt.Println(" Func1 is executed")
}
func Func2(Name string) int { // this func will return int type here
	return 1900
}