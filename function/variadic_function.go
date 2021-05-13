package main

import "fmt"

func main() {
	CalculateMarks("bhanvi")
	CalculateMarks("bachhaniya",12,45)

}

func CalculateMarks(Name string,Mark ...int){
//	fmt.Println(Mark)
//	fmt.Println(Name)
for _, value := range Mark  {
	fmt.Println(Name,value)
}
fmt.Println("====")
}