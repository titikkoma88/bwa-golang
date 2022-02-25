package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{
		ID:        1,
		FirstName: "Waqqosh",
		LastName:  "Hadjarati",
		Email:     "mail@waqqosh.com",
		IsActive:  true,
	}

	user2 := User{2, "panji", "Hadjarati", "mail@panjihadjarati.com", true}

	users := []User{user, user2}

	group := Group{"Programmer", user, users, true}

	displayGroup(group)
}

func displayGroup(group Group)  {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
