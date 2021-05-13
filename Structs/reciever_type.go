package main
import "fmt"

type Person struct{
	Fname string
	LName string

}
//value type receiver
func(per Person) SetPerDetails(){
	per.Fname="ererer"

}

//refe TR
func(per *Person) SetPerDetailss(){
	per.Fname="ererer1"

}



func main(){
per:=Person{"rtr","er"}
per.SetPerDetails()
per.SetPerDetailss()
fmt.Println(per)
}