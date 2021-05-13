package main

import (
	"fmt" //IN BUILD PACAKGE
	"golang.org/x/example/stringutil" //THIRD PARTY PACKGAE
	"rsc.io/quote"

	"myProject.com/DAL" //MODULENAME +PATH
)


func main(){
	fmt.Println(stringutil.Reverse("abc"))
	fmt.Println(quote.Hello())
	DAL.ReadDB()
	DAL.WriteDB()
	DAL.ReadSQLDB()
	DAL.WriteSQLDB()
}