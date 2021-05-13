package main
import "fmt"
func main() {
	//Interpreted String
	// var str1 = "This is First Day of GoLang.
	// This is basic session
	// "
	//Raw String
	// var str2 = `This is First Day of Golang
	// This is basic session
	
	// `
	//fmt.Println(str2)
	var str1 = "This is First Day's of GoLang. \n  This is basic session"
	fmt.Println(str1)
	var str2 = `This is First "Day" of GoLang. \n  This is basic session`
	fmt.Println(str2)
}
