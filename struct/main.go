package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Panji"
	user.LastName = "Hadjarati"
	user.Email = "mail@panjihadjarati.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Panji"
	user2.LastName = "Ramadhan"
	user2.Email = "jie.piranha@gmail.com"
	user2.IsActive = true

	user3 := User{
		ID:        3,
		FirstName: "Waqqosh",
		LastName:  "Hadjarati",
		Email:     "mail@waqqosh.com",
		IsActive:  true,
	}

	user4 := User{4, "Shofiyyah", "Hadjarati", "mail@shofiyyah.com", true}

	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)
	fmt.Println(user4)

	fmt.Println("===========")

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user4)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}
