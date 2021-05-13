package main

import "fmt"

func main(){

	weekdays := []string{"mon","tue","wed","thr","fri"}

	weekdays =append(weekdays[:3],weekdays[5:]...)
	fmt.Println(weekdays)
}