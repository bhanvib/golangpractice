package main
import "fmt"
func main() {
	//var slice1 = []string{"GoLang", "JavaScript"} //slice
	//var slice1 []int //nil slice
	//fmt.Println(slice1)
	//fmt.Println(slice1[0])
	//slice1 = append(slice1, 40,60)
	var slice1 = make([]int,2)  // length and capacity is 2
	fmt.Println("Len ",len(slice1))  //2
	fmt.Println(slice1)
	slice1= append(slice1, 30,40)
	fmt.Println("Len ",len(slice1)) //4
	fmt.Println(slice1)
}