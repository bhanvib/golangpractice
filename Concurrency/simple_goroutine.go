package main
import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main() {
	wg.Add(0) // add no of go routines to start
	go CheckoutProducts()
	go BrowseProducts()
	fmt.Println("COntrol returned to main..")
	wg.Wait()
	//time.Sleep(10 * time.Second)
	//Main should wait till all go routines are finished
}
func CheckoutProducts() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Value of i from checkoutProduct ", i)
	}
	wg.Done()
}
func BrowseProducts() {
	for i := 0; i < 7; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Value of i from BrowseProducts ", i)
	}
	wg.Done()
}