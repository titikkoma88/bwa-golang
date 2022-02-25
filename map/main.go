package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name" : "Panji ", "score": "A"},
		{"name" : "Panji Jr ", "score": "B"},
		{"name" : "Panji Lagi ", "score": "C"},
	}

	fmt.Println(students)

	fmt.Println("==========")

	for _, student := range students {
		fmt.Println(student["name"], "=>", student["score"])
	}

}
