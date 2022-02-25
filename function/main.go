package main

import "fmt"

func main()  {
	result := add(10, 20)
	fmt.Println(result)
}

func add(number, numberTwo int) int {
	return number + numberTwo
	
}

