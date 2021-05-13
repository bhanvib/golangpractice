package main

import"fmt"

func main(){
	var i = 100

	var iPointer *int
	iPointer =&i

	fmt.Println(iPointer)
	fmt.Println("before calling", i)
	changeValue(&i) //passing pointer to i
	fmt.Println("before calling" ,i)
	
	
	array2 := [2]int{100, 200}
    changeArray(array2)
	fmt.Println(array2)
	

}

func changeValue(i *int){
  *i=1000 //access value from ipointet (mem address)
}

func changeArray(array1 [2] int){
	array1[0]=3030303

}

//slice will pass refrance
