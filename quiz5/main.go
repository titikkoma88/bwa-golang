package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name string
	Admin User
	Users []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func main() {
	user := User{
		ID:        1,
		FirstName: "Waqqosh",
		LastName:  "Hadjarati",
		Email:     "mail@waqqosh.com",
		IsActive:  true,
	}

	result := user.Display()

	fmt.Println(result)

	user2 := User{2, "panji", "Hadjarati", "mail@panjihadjarati.com", true}
	fmt.Println(user2.Display())
	
	users := []User{user, user2}

	group := Group{"Programmer", user, users, true}

	group.DisplayGroup()
}

