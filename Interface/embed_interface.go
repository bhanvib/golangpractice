package main()

import "fmt"
type interface1 interface{
	Test1()
}
type interface2 interface{
	interface1
	Test2()
} 
func main(){

}