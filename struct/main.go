package main

import (
	"fmt"
	"struct/management"
)

func main() {
	user := management.User{
		ID:        1,
		FirstName: "Waqqosh",
		LastName:  "Hadjarati",
		Email:     "mail@waqqosh.com",
		IsActive:  true,
	}

	result := user.Display()

	fmt.Println(result)

	user2 := management.User{2, "panji", "Hadjarati", "mail@panjihadjarati.com", true}
	fmt.Println(user2.Display())
	
	users :=[]management.User{user, user2}

	group := management.Group{"Programmer", user, users, true}

	group.DisplayGroup()

}

