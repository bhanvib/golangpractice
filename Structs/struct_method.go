package main
import "fmt"
type Item struct {
	ItemId   int
	ItemName string
}
//Method : FUnction with receiver Type
func (item Item) GetItem() {
	fmt.Println("Get Item is called for ", item.ItemName)
}
//FUnction
// func GetItem(){
// }
func main() {
	//GetItem()
	item1 := Item{ItemId: 10, ItemName: "TV"}
	item1.GetItem()
}

