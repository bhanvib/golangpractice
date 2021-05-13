package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main() {
	ch1 := make(chan int)
	wg.Add(2) // add no of go routines to start
	go ReadDB(ch1)
	go WriteFile(ch1)
	fmt.Println("COntrol returned to main..")
	wg.Wait()
}
func ReadDB(ch1 chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Put value into channel")
		ch1 <- i // Writing into channel
	}
	wg.Done()
}
func WriteFile(ch1 chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Read From Channel", <-ch1) //Read From channel
	}
	wg.Done()
}
