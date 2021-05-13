package main

import "fmt"

func main(){

	var array1 [3]int

	
	array1[0]=100
	array1[1]=200
	array1[2]=300
	

	for i:=0;i<len(array1);i++{
		fmt.Println(array1[i])
	}
	for i,v:= range array1{
		fmt.Println(i,v)
	}
	var array2 =[2]string{"rerr","rtet"}
	fmt.Println(array2)
}