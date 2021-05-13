package main
import (
	"fmt"
	"time"
)
func Server11(ch1 chan string) {
	time.Sleep(3 * time.Second)
	ch1 <- "Response from Server1"
}
func Server22(ch2 chan string) {
	time.Sleep(6 * time.Second)
	ch2 <- "Response from Server2"
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Server11(ch1)
	go Server22(ch2)
	merge(ch1, ch2)
}
func merge(data ...chan string) { // slice of channels
	for _, ch1 := range data {
		fmt.Println(<-ch1)
	}
}