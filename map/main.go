package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["javaScript"] = 9
	myMap["go"] = 9

	fmt.Println(myMap)

	languageMap := map[string]string{
		"ruby": "is beatifull",
		"go":   "is super fast,",
		"javaScript": "hype",
	}

	for key, value := range languageMap {
		fmt.Println("Key : ", key, " value : ", value)
	}

	fmt.Println("==========")

	delete(languageMap, "go")

	for key, value := range languageMap {
		fmt.Println("Key : ", key, " value : ", value)
	}


	value, isAvailable := languageMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)

}
