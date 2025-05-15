package main
import "fmt"

type User struct {
	name string
	email string
}

func (u User) GetInfo(){
	fmt.Println("USER INFO")
	fmt.Println("=========")
	fmt.Printf("Name : %s\n", u.name)
	fmt.Printf("Email : %s\n", u.email)
}

func (u *User) SetName(newName string){
	*&u.name = newName  
}

func main() {
	
	user := User{name: "Ambarasudi", email: "ambarasudi@gmail.com"}
	user.GetInfo()

	user.SetName("Labubu")
	user.GetInfo()

}