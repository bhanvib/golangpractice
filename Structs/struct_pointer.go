package main
import "fmt"

type Customer struct{
	CustomerId int
	Address string
	Name string

}
func main(){
	cust1:=Customer{10,"pune","fj"}
	//cust1 :=&Customer{10,"pune","fj"} way2 
	changeStruct(&cust1)
	fmt.Println(cust1)
}
func changeStruct(cust *Customer){
	cust.Name="dfgs"
}