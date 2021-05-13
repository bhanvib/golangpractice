package main
import "fmt"
func main() {

	//var array2 =[2]string{"rerr","rtet"}//array

	var slice  = []string{"rerr","rtet"}

	for i,v:=range slice{

		fmt.Println(i,v)
	}
	slice=append(slice,"njdfje","deer")
	fmt.Println((slice))

}