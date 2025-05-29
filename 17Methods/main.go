package main

import "fmt"

func main() {
	parth := User{"parth", "gg@gg.ai", true, 21}
	fmt.Println(parth)

	parth.GetStatus()
	parth.changeEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) changeEmail() {
	u.Email = "gg@gg.com"
	fmt.Println("User's mail id: ", u.Email)
}
