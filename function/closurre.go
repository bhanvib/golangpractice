package main
import "fmt"
//var counter = 3 //global
// Counter should be local
//still it should be initialized only once
//From outer func we should get inner func which will increment this one time initialized counter
// func AddHits() int {
// 	var counter = 0
// 	counter++
// 	return counter
// }
var getHits = func() func() int {
	var counter = 0
	var addHits = func() int {
		counter++
		return counter
	}
	return addHits // returning defn of inner/closure func
}() // after execution it is giving us inner func defn which will have access to counter
func main() {

	fmt.Println(getHits()) // calling addHits
	fmt.Println(getHits()) // calling addHits
	fmt.Println(getHits()) // calling addHits
	//fmt.Println(getHits()) //get defn of closure /inner func
	// fmt.Println(getHits()())
	// fmt.Println(getHits()())
	// fmt.Println(AddHits())
	// fmt.Println(AddHits())
	// //counter = 190
	// fmt.Println(AddHits())
}