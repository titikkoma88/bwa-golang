package main

import "fmt"

func main() {

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Saya sedang belajar Go:", i)
	// }

	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Saya sedang belajar Go:", i)
	// 	i++
	// }

	title := "Golang the best language"

	// for index, letter := range title {
	// 	// fmt.Println("index: ", index, " letter :", letter)
	// 	fmt.Println("index: ", index, " letter :", string(letter))

	// }

	for _, letter := range title {
		fmt.Println("letter :", string(letter))

	}

}
