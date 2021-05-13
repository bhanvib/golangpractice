package main
import (
	"fmt"
	"strconv"
)
func main() {
	var str1 = "1234ab"
	value, _ := strconv.Atoi(str1) // convert from string to int
	fmt.Println(value)
	strValue := strconv.Itoa(123)
	fmt.Println(strValue)
}