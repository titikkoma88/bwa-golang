package main

import "fmt"

func main()  {
	printMyResult()
	fmt.Println("==========")
	printMyResult2("Saya sedang")
	printMyResult2("belajar")
	printMyResult2("Golang Fundamental")
	fmt.Println("==========")
	sentence2 := printMyResult3("Materi")
	fmt.Println(sentence2)

}

// 1. Input
// 2. Proses
// 3. Output

func printMyResult()  {
	fmt.Println("Saya sedang belajar golang")
}

func printMyResult2(sentence string)  {
	fmt.Println(sentence)
}

func printMyResult3(sentence2 string) string {
	newSentence := sentence2 + " function"
	return newSentence
}
