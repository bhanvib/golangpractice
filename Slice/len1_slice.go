package main
import "fmt"
func main() {
	//len: current no of elements in internal array
	//cap : cap of internal array to hold no of elements
	var slice2 = make([]int, 2, 4)
	fmt.Println(slice2)
	fmt.Println("Len", len(slice2), "cap ", cap(slice2))
	slice3 := append(slice2, 40, 50, 70) // 4 : Both slice2 and slice3 are refering to same internal array
	fmt.Println(slice3)
	slice3[0] = 888
	fmt.Println(slice3)
	fmt.Println(slice2)
	// fmt.Println("Len", len(slice2), "cap ", cap(slice2)) // it was under/= cap
	// slice2 = append(slice2, 70) // now cap will exceed
	// fmt.Println(slice2)
	// fmt.Println("Len", len(slice2), "cap ", cap(slice2))
}