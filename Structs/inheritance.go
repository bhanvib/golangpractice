package main
import "fmt"
type User struct { // Parent/Base
	UserId   string
	Password string
}
func (user User) Login() {
	fmt.Println(" User is logged in...")
}
type Admin struct {
	//Embed User struct here
	User // will reuse Memebrs of User struct here
	Role string
}
func (admin Admin) Login(){
	fmt.Println("efer")
}
func main() {
	user1 := User{"admin", "1234"}
	admin1 := Admin{user1, "Sytem Admin"}
	fmt.Println(admin1.UserId, admin1.Password)
	admin1.Login()
}