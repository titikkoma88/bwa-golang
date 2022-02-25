package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func main() {
	user := User{
		ID:        1,
		FirstName: "Waqqosh",
		LastName:  "Hadjarati",
		Email:     "mail@waqqosh.com",
		IsActive:  true,
	}

	result := user.display()

	fmt.Println(result)

	user2 := User{2, "panji", "Hadjarati", "mail@panjihadjarati.com", true}
	fmt.Println(user2.display())
	
}

