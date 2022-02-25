package main

import "fmt"

func main()  {
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["javaScript"] = 9
	myMap["go"] = 9

	fmt.Println(myMap)

	languageMap := map[string]string{
		"ruby": "is beatifull",
		"go": "is super fast,",
	}

	fmt.Println(languageMap)
}