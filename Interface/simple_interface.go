package main
import "fmt"
type IShape interface {
	Area()
	Perimeter()
}
type Triangle struct {
	side int
}
func (tra Triangle) Area() {
	fmt.Println("Area is calculated..")
}
func (tra Triangle) Perimeter() {
	fmt.Println("Perimeter is calculated..")
}
type Square struct {
	side int
}
func (sqr Square) Area() {
	fmt.Println("Area is calculated for Square..")
}
func (sqr Square) Perimeter() {
	fmt.Println("Perimeter is calculated for Square..")
}
func main() {
	var shape IShape
	var shape1 IShape
	shape = Triangle{10}
	shape.Area()
	shape.Perimeter()
	shape1 = Square{20}
	shape1.Area()
	shape1.Perimeter()
}

