package main

import "fmt"

func main()  {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	number := 5
	fmt.Println("Alamat memory :", &number)
	fmt.Println("Nilai awal :", number)

	change(number, 100)
	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memory :", &number)

	fmt.Println("=========")

	change2(&number, 100)
	fmt.Println("Nilai akhir 2 :", number)
	fmt.Println("Alamat memory 2 :", &number)
}

func change(old int, new int)  {
	old = new
	fmt.Println("Alamat memory :", &old)
	fmt.Println("Didalam function :", old)
}

func change2(old *int, new int)  {
	*old = new
	fmt.Println("Alamat memory 2 :", old)
	fmt.Println("Didalam function 2 :", *old)
}