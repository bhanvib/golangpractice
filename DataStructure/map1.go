package main

import "fmt"

func main(){

//first way

	var map1=map[string]int{"key1":10,"key2":24,"key3":35}

	for key1,value1 :=range map1 {
		fmt.Println(key1,value1)
	}

	//2way

	var map2 =make(map[int]string)
	map2[10]="d"
	map2[20]="34"

	fmt.Println(map2[20])
	map2[20] = "ASAP.NET MVC" //update
	delete(map2,20)
	fmt.Println(map2)

	value,status :=map2[20]
	fmt.Println(value,status)

	if value1, exist := map2[300]; exist {
		fmt.Println("Key 300 exists", value1)
	} else {
		fmt.Println("Key 30 does nor exists")
	}
}