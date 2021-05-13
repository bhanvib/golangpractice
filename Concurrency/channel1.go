package main
import (
	"fmt"
	"time"
)
func BrowseOtherProducts() {
	fmt.Println("Browse Products..")
}
func CheckoutCart(ch1 chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Checkout Product is done")
	ch1 <- 1900
}
func DeliverProducts(price int) {
	fmt.Println("Deliver Products..")
}
func main() {
	ch1 := make(chan int)
	go CheckoutCart(ch1)
	fmt.Println("Control returned to main method")
	BrowseOtherProducts() // this will not wait till Checkout is completed
	//Read Data from ch1
	price := <-ch1         // Waiting operation
	DeliverProducts(price) // it should wait till checkout is completed
	//time.Sleep(10 * time.Second)
}

//wg - for waiting the function and channel will use for transfering the data 