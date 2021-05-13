package main
import "fmt"

func main() {
	ch := make(chan int,4)
	ch <- 100
	ch <- 200
	fmt.Println("gdggd", <-ch)

	ch <- 400
	fmt.Println("gdggd", <-ch)
}

